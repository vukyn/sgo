package analyzer

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/vukyn/kuery/query"
)

var progressMessages = []string{
	"Investigating your codebase...",
	"Scanning your codebase...",
	"Analyzing your code...",
	"Uncovering your secrets...",
	"Predicting your code's future...",
	"Casting analysis spells...",
	"Calculating code metrics...",
	"Uncovering plot twists in your code...",
	"Messing with your code files...",
}

func getRandomProgressMessage() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return progressMessages[r.Intn(len(progressMessages))]
}

func formatList(items []string) string {
	if len(items) == 0 {
		return "None"
	}
	return strings.Join(items, "\n")
}

func formatPackages(packages map[string]string) string {
	if len(packages) == 0 {
		return "None"
	}
	var result []string
	for pkg, version := range packages {
		result = append(result, fmt.Sprintf("%s: %s", pkg, version))
	}
	return strings.Join(result, "\n")
}

func formatFrameworks(frameworks map[string]struct{}) string {
	if len(frameworks) == 0 {
		return "None"
	}
	return strings.Join(query.Keys(frameworks), "\n")
}

func formatWarnings(warnings []string) string {
	if len(warnings) == 0 {
		return "None"
	}
	return strings.Join(warnings, "\n")
}

func formatSecretFindings(findings []SecretFinding) string {
	if len(findings) == 0 {
		return "None"
	}
	var result []string
	for _, finding := range findings {
		result = append(result, fmt.Sprintf("[%s] %s\n  %s\n  %s",
			finding.Category,
			finding.Description,
			finding.File,
			finding.Line))
	}
	return strings.Join(result, "\n")
}

func formatSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}
