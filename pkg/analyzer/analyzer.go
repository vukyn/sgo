package analyzer

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/schollz/progressbar/v3"
	"golang.org/x/mod/modfile"
)

type AnalysisResult struct {
	ScanDuration    time.Duration       `json:"scan_duration"`
	GoVersion       string              `json:"go_version"`
	Frameworks      map[string]struct{} `json:"frameworks"`
	SecretKeys      []SecretFinding     `json:"secret_keys"`
	TODOs           []string            `json:"todos"`
	EmptyGoFiles    []string            `json:"empty_go_files"`
	EmptyOtherFiles []string            `json:"empty_other_files"`
	TotalLines      int                 `json:"total_lines"`
	CommentLines    int                 `json:"comment_lines"`
	EmptyLines      int                 `json:"empty_lines"`
	Packages        map[string]string   `json:"packages"` // map[package_name]version
	ProjectSize     int64               `json:"project_size"`
	TotalGoFiles    int                 `json:"total_go_files"`
	Warnings        []string            `json:"warnings"`
	IgnoredPatterns []string            `json:"ignored_patterns"`
	Summary         SummaryStatus       `json:"summary"`
}

type SecretFinding struct {
	File        string `json:"file"`
	Line        string `json:"line"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

type fileResult struct {
	path     string
	content  string
	isGoFile bool
}

type Analyzer struct {
	path            string
	ignoredPatterns []string
	progressBar     *progressbar.ProgressBar
}

func NewAnalyzer(path string) *Analyzer {
	return &Analyzer{
		path:            path,
		ignoredPatterns: DefaultIgnoredPatterns,
	}
}

func (a *Analyzer) Analyze() (*AnalysisResult, error) {
	now := time.Now()
	result := &AnalysisResult{
		Packages:        make(map[string]string),
		Frameworks:      make(map[string]struct{}),
		IgnoredPatterns: a.ignoredPatterns,
	}

	// Initialize progress bar with random message
	a.progressBar = progressbar.Default(-1, getRandomProgressMessage())

	// Analyze go.mod file
	if err := a.analyzeGoMod(result); err != nil {
		return nil, err
	}

	// Create channels for the fan-in fan-out pattern
	fileChan := make(chan fileResult, 100)
	pathChan := make(chan string, 100)
	doneChan := make(chan struct{})

	// Get number of CPU cores for optimal worker count
	numWorkers := runtime.NumCPU()
	var wg sync.WaitGroup

	// Start file scanner
	go func() {
		defer close(pathChan)
		a.scanFiles(pathChan)
	}()

	// Start worker goroutines
	for range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range pathChan {
				content, err := os.ReadFile(path)
				if err != nil {
					continue
				}

				isGoFile := strings.HasSuffix(path, ".go")
				fileChan <- fileResult{
					path:     path,
					content:  string(content),
					isGoFile: isGoFile,
				}
				a.progressBar.Add(1)
			}
		}()
	}

	// Start result processor
	go func() {
		for file := range fileChan {
			if file.isGoFile {
				result.TotalGoFiles++
				a.analyzeGoFile(file.path, file.content, result)
			}
			a.analyzeFile(file.path, result)
		}
		doneChan <- struct{}{}
	}()

	// Wait for all workers to finish and close channels
	go func() {
		wg.Wait()
		close(fileChan)
	}()

	// Wait for result processing to complete
	<-doneChan

	if len(result.SecretKeys) > 0 {
		result.Warnings = append(result.Warnings, "*potential secret key found in codebase")
	}
	result.ScanDuration = time.Since(now)
	result.SummaryStatus()
	return result, nil
}

func (a *Analyzer) scanFiles(pathChan chan<- string) {
	err := filepath.Walk(a.path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		for _, pattern := range a.ignoredPatterns {
			if matched, _ := filepath.Match(pattern, info.Name()); matched {
				if info.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}

		if !info.IsDir() {
			pathChan <- path
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error scanning files: %v\n", err)
	}
}

func (a *Analyzer) analyzeGoMod(result *AnalysisResult) error {
	modPath := filepath.Join(a.path, "go.mod")
	if _, err := os.Stat(modPath); os.IsNotExist(err) {
		result.Warnings = append(result.Warnings, "*go.mod file not found")
		return nil
	}

	content, err := os.ReadFile(modPath)
	if err != nil {
		return err
	}

	f, err := modfile.Parse("go.mod", content, nil)
	if err != nil {
		return err
	}

	if f.Go != nil {
		result.GoVersion = f.Go.Version
	}

	for _, req := range f.Require {
		result.Packages[req.Mod.Path] = req.Mod.Version
	}

	return nil
}

func (a *Analyzer) analyzeGoFile(path, content string, result *AnalysisResult) {
	lines := strings.Split(content, "\n")

	// Check for empty files
	switch len(lines) {
	case 0, 1:
		hasPackage := false
		for _, line := range lines {
			if strings.HasPrefix(strings.TrimSpace(line), "package ") {
				hasPackage = true
				break
			}
		}
		if hasPackage {
			result.EmptyGoFiles = append(result.EmptyGoFiles, path) // empty go file
		} else {
			result.EmptyOtherFiles = append(result.EmptyOtherFiles, path) // empty other file
		}
	default:
		for lineNum, line := range lines {
			line = strings.TrimSpace(line)

			if line == "" {
				result.EmptyLines++
				continue
			}

			if strings.HasPrefix(line, "//") {
				result.CommentLines++
				if strings.Contains(strings.ToLower(line), "todo:") {
					result.TODOs = append(result.TODOs, fmt.Sprintf("%s:%d: %s", path, lineNum+1, strings.TrimSpace(line)))
				}
				continue
			}

			result.TotalLines++

			// Check for frameworks (ignore struct fields and comments)
			if !strings.Contains(line, "struct") && !strings.HasPrefix(line, "//") {
				if strings.Contains(line, "gin.") {
					result.Frameworks["gin"] = struct{}{}
				}
				if strings.Contains(line, "echo.") {
					result.Frameworks["echo"] = struct{}{}
				}
				if strings.Contains(line, "fiber.") {
					result.Frameworks["fiber"] = struct{}{}
				}
			}

			// Check for secret keys (ignore struct fields and comments)
			if !strings.Contains(line, "struct") && !strings.HasPrefix(line, "//") {
				for _, pattern := range DefaultSecretPatterns {
					if matched, _ := regexp.MatchString(pattern.Pattern, line); matched {
						result.SecretKeys = append(result.SecretKeys, SecretFinding{
							File:        path,
							Line:        fmt.Sprintf("%d: %s", lineNum+1, strings.TrimSpace(line)),
							Description: pattern.Description,
							Category:    pattern.Category,
						})
						break
					}
				}
			}
		}
	}
}

func (a *Analyzer) analyzeFile(path string, result *AnalysisResult) {
	info, err := os.Stat(path)
	if err != nil {
		return
	}
	result.ProjectSize += info.Size()
}

func (r *AnalysisResult) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder.Encode(r)
}

func (r *AnalysisResult) ToText(w io.Writer) error {
	_, err := fmt.Fprintf(w, `Project Analysis Results:
--------------------------------
Status: %s (Score: %d)
Notes:
%s

Scan Duration: %s
Go Version: %s
Total Go Files: %d
Total Lines: %d
Comment Lines: %d
Empty Lines: %d
Project Size: %s

Frameworks Used:
%s

TODOs Found:
%s

Empty Go Files:
%s

Empty Other Files:
%s

Potential Secret Keys:
%s

Packages Used:
%s

Warnings:
%s
`,
		r.Summary.Status,
		r.Summary.Score,
		formatList(r.Summary.Notes),
		r.ScanDuration,
		r.GoVersion,
		r.TotalGoFiles,
		r.TotalLines,
		r.CommentLines,
		r.EmptyLines,
		formatSize(r.ProjectSize),
		formatFrameworks(r.Frameworks),
		formatList(r.TODOs),
		formatList(r.EmptyGoFiles),
		formatList(r.EmptyOtherFiles),
		formatSecretFindings(r.SecretKeys),
		formatPackages(r.Packages),
		formatWarnings(r.Warnings),
	)
	return err
}
