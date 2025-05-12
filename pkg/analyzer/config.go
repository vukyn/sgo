package analyzer

const (
	// Score system
	BaseScore                            = 100
	LargeProjectSizeThreshold            = 100 * 1024 * 1024 // 100MB
	DeductionLargeProjectSize            = 5
	HighNumberOfTODOsThreshold           = 15
	DeductionHighNumberOfTODOs           = 15
	MultiplePotentialSecretKeysThreshold = 5
	DeductionMultiplePotentialSecretKeys = 20
	MultipleEmptyFilesThreshold          = 5
	DeductionMultipleEmptyFiles          = 10
	MultipleWarningsThreshold            = 5
	DeductionMultipleWarnings            = 15
)

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
