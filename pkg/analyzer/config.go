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

type SecretPattern struct {
	Pattern     string
	Description string
	Category    string
}

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

	DefaultSecretPatterns = []SecretPattern{
		// API Keys and Tokens
		{Pattern: `(?i)(api[_-]?key|apikey)["']?\s*[:=]\s*["']?[a-zA-Z0-9]{32,}`, Description: "API Key", Category: "API"},
		{Pattern: `(?i)(jwt[_-]?token|jwt)["']?\s*[:=]\s*["']?[a-zA-Z0-9-_=]+\.[a-zA-Z0-9-_=]+\.?[a-zA-Z0-9-_.+/=]*`, Description: "JWT Token", Category: "Token"},
		{Pattern: `(?i)(access[_-]?token|access_token)["']?\s*[:=]\s*["']?[a-zA-Z0-9]{32,}`, Description: "Access Token", Category: "Token"},
		{Pattern: `(?i)(refresh[_-]?token|refresh_token)["']?\s*[:=]\s*["']?[a-zA-Z0-9]{32,}`, Description: "Refresh Token", Category: "Token"},

		// Certificates and Keys
		{Pattern: `(?i)-----BEGIN\s+(?:RSA\s+)?PRIVATE\s+KEY-----`, Description: "Private Key", Category: "Certificate"},
		{Pattern: `(?i)-----BEGIN\s+CERTIFICATE-----`, Description: "Certificate", Category: "Certificate"},
		{Pattern: `(?i)\.(pem|key|crt|cer|der|p12|pfx)$`, Description: "Certificate File", Category: "Certificate"},

		// Database Credentials
		{Pattern: `(?i)(db[_-]?(?:password|pass|pwd))["']?\s*[:=]\s*["']?[^"'\s]+`, Description: "Database Password", Category: "Database"},
		{Pattern: `(?i)(mongodb[_-]?uri|mongo[_-]?uri)["']?\s*[:=]\s*["']?mongodb(\+srv)?://[^"'\s]+`, Description: "MongoDB URI", Category: "Database"},
		{Pattern: `(?i)(postgres[_-]?uri|pg[_-]?uri)["']?\s*[:=]\s*["']?postgres(ql)?://[^"'\s]+`, Description: "PostgreSQL URI", Category: "Database"},

		// Cloud Credentials
		{Pattern: `(?i)(aws[_-]?(?:access[_-]?key|secret[_-]?key|secret))["']?\s*[:=]\s*["']?[A-Z0-9]{20,}`, Description: "AWS Credentials", Category: "Cloud"},
		{Pattern: `(?i)(gcp[_-]?(?:key|credentials|secret))["']?\s*[:=]\s*["']?[A-Za-z0-9+/]{32,}`, Description: "GCP Credentials", Category: "Cloud"},
		{Pattern: `(?i)(azure[_-]?(?:key|secret|connection[_-]?string))["']?\s*[:=]\s*["']?[A-Za-z0-9+/]{32,}`, Description: "Azure Credentials", Category: "Cloud"},

		// General Secrets
		{Pattern: `(?i)(password|passwd|pwd)["']?\s*[:=]\s*["']?[^"'\s]+`, Description: "Password", Category: "General"},
		{Pattern: `(?i)(secret[_-]?key|secret)["']?\s*[:=]\s*["']?[a-zA-Z0-9]{32,}`, Description: "Secret Key", Category: "General"},
		{Pattern: `(?i)(auth[_-]?token|auth_token)["']?\s*[:=]\s*["']?[a-zA-Z0-9]{32,}`, Description: "Auth Token", Category: "General"},
	}
)
