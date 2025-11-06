# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Purpose

This is a Go learning repository containing demos for studying various Go concepts including goroutines, channels, concurrency patterns, and other language features.

## Repository Structure

The repository uses a shared `go.mod` at the root level. Demos are organized as:

- **Simple examples** - Single `.go` files in the root directory (e.g., `goroutine_basic.go`, `channel_demo.go`)
- **Complex examples** - Subdirectories with multiple files (e.g., `web-server/`, `worker-pool/`)

This structure allows:
- Quick single-file demos that can be run directly
- Organized multi-file projects in their own folders
- Shared dependencies through the root `go.mod`

## Running Demos

### Running a single-file demo:
```bash
go run goroutine_basic.go
```

### Running a folder-based demo:
```bash
go run ./web-server
# or
cd web-server && go run .
```

### Running with race detection (important for concurrency demos):
```bash
go run -race goroutine_basic.go
# or
go run -race ./worker-pool
```

### Building a demo:
```bash
go build -o myapp ./demo-folder
./myapp
```

## Creating New Demos

### For simple examples (single concept):
Create a `.go` file directly in the root:
```bash
# File: channel_basics.go
package main

func main() {
    // demo code here
}
```

Run with: `go run channel_basics.go`

### For complex examples (multiple files/packages):
Create a subdirectory:
```bash
mkdir worker-pool
cd worker-pool
```

Create files with `package main` and `main()` function.

Run with: `go run .` or `go run ./worker-pool` from root

### When to use a separate folder:
- Multiple source files
- Need internal packages
- Larger, more complex examples
- Need separate configuration or resources

## Go Commands Reference

- `go run .` - Compile and run the current package
- `go run -race .` - Run with race detector (use for concurrency demos)
- `go build` - Compile the package
- `go fmt ./...` - Format all Go files
- `go vet ./...` - Report suspicious constructs
- `go test ./...` - Run tests if present

## Working with Concurrency Demos

When working with goroutines and channels:
- Always test with `-race` flag to detect race conditions
- Use proper synchronization (channels, sync.WaitGroup, sync.Mutex, etc.)
- Demonstrate both correct and incorrect patterns when educational
- Include comments explaining the concurrency behavior

## Code Organization

Each demo should:
- Be self-contained and executable
- Include clear comments explaining the concept being demonstrated
- Have a descriptive folder name (e.g., `goroutine-basics`, `channel-select`, `worker-pool`)
- Focus on a single concept or pattern for clarity
