# SGO - Go Project Analyzer

SGO is a command-line tool for analyzing and visualizing Go project structure and codebase. It provides detailed insights about your Go projects, including code metrics, dependencies, and potential issues.

## Features

-   [x] Scan Go files and relevant project files (go.mod, etc.)
-   [x] Detect Go version
-   [x] Identify used Go frameworks
-   [x] Detect potential secret keys in codebase
-   [x] Count TODO comments
-   [x] Count total .go files
-   [x] Count total lines of code, comments, and empty lines
-   [x] Count total packages used (direct and indirect)
-   [x] Identify empty Go files and other files
-   [x] Calculate total project size
-   [x] Ignore cache/config folders and files
-   [x] Fast and efficient file scanning
-   [x] Output in JSON or text format
-   [x] Summary status of the project

## Installation

```bash
go install github.com/vukyn/sgo@latest
```

## Usage

```bash
# Analyze current directory
sgo

# Analyze specific directory
sgo -p /path/to/project

# Output in text format (default)
sgo -o text

# Output in JSON format
sgo -o json
```

## Output Format

### JSON Output

```json
{
	"scan_duration": "1.23s",
	"go_version": "1.21",
	"frameworks": ["gin", "echo"],
	"secret_keys": [],
	"todos": ["// TODO: Implement error handling"],
	"empty_go_files": [],
	"empty_other_files": [],
	"total_lines": 1000,
	"comment_lines": 200,
	"empty_lines": 100,
	"packages": {
		"github.com/gin-gonic/gin": "v1.9.1"
	},
	"project_size": 1024000,
	"total_go_files": 10,
	"warnings": ["*potential secret key found in codebase"],
	"ignored_patterns": [".git", "vendor", ".vscode"],
	"summary": {
		"status": "BAD",
		"score": 75,
		"notes": ["Large project size (> 100MB)", "Has 15 TODOs"]
	}
}
```

### Text Output

```bash
Project Analysis Results:
--------------------------------
Status: BAD (Score: 75)
Notes:
- Large project size (> 100MB)
- Has 15 TODOs

Scan Duration: 1.23s
Go Version: 1.21
Total Go Files: 10
Total Lines: 1000
Comment Lines: 200
Empty Lines: 100
Project Size: 1024000 bytes

Frameworks Used:
gin
echo

TODOs Found:
// TODO: Implement error handling

Empty Go Files:
None

Empty Other Files:
None

Potential Secret Keys:
None

Packages Used:
github.com/gin-gonic/gin: v1.9.1

Warnings:
*potential secret key found in codebase
```

## License

MIT License
