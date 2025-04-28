# SGO - Go Project Analyzer

SGO is a command-line tool for analyzing and visualizing Go project structure and codebase. It provides detailed insights about your Go projects, including code metrics, dependencies, and potential issues.

## Features

- Scan Go files and relevant project files (go.mod, etc.)
- Detect Go version
- Identify used Go frameworks
- Detect potential secret keys in codebase
- Find TODO comments
- Identify empty Go files
- Count total lines of code, comments, and empty lines
- Count total packages used (direct and indirect)
- Calculate total project size
- Count total .go files
- Ignore cache/config folders and files
- Show progress bar while scanning
- Concurrent file scanning using goroutines
- Output in JSON or text format

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

# Output in text format
sgo -o text

# Output in JSON format (default)
sgo -o json
```

## Output Format

### JSON Output

```json
{
 "go_version": "1.21",
 "frameworks": ["gin", "echo"],
 "secret_keys": [],
 "todos": ["// TODO: Implement error handling"],
 "empty_files": [],
 "total_lines": 1000,
 "comment_lines": 200,
 "empty_lines": 100,
 "packages": {
  "github.com/gin-gonic/gin": "v1.9.1"
 },
 "project_size": 1024000,
 "total_go_files": 10,
 "ignored_patterns": [".git", "vendor", ".vscode"]
}
```

### Text Output

```bash
Project Analysis Results:
-------------------
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

Empty Files:
None

Potential Secret Keys:
None

Packages Used:
github.com/gin-gonic/gin: v1.9.1
```

## License

MIT License
