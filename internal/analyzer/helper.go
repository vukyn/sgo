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
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // #nosec G404 -- cosmetic RNG for progress-bar message, no security impact
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
