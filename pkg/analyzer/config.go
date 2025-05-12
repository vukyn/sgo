package analyzer

var (
	DefaultIgnoredPatterns = []string{
		".git",
		"vendor",
		".vscode",
		"node_modules",
		"*.exe",
		"*.test",
		"*.out",
	}

	DefaultSecretKeys = []string{
		"apikey",
		"secret",
		"password",
		"token",
	}
)
