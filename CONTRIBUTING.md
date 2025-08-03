# Contributing to Tea Timer

Thank you for your interest in contributing to Tea Timer! This document provides guidelines and information for contributors.

## Getting Started

### Prerequisites

- Go 1.24 or later
- Git
- Terminal with color support (for testing)

### Setting Up the Development Environment

1. Fork the repository
2. Clone your fork:
   ```bash
   git clone https://github.com/your-username/tea-timer.git
   cd tea-timer
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Build the project:
   ```bash
   go build
   ```

## Development Guidelines

### Code Style

- Follow Go's official [Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use `gofmt` to format your code
- Keep functions small and focused
- Add comments for complex logic
- Use meaningful variable and function names

### Testing

Run tests before submitting a pull request:
```bash
go test ./...
```

### Building

Test your changes by building the project:
```bash
go build
go run main.go 10  # Test with 10-second timer
```

## Making Changes

### Branch Naming

Use descriptive branch names:
- `feature/add-sound-notification`
- `bugfix/fix-timer-pause-issue`
- `docs/update-readme-installation`

### Commit Messages

Follow conventional commit format:
```
type(scope): description

feat: add sound notification when timer finishes
fix: resolve timer pause behavior
docs: update installation instructions
style: format code with gofmt
refactor: simplify timer logic
```

### Pull Request Process

1. Create a feature branch from `main`
2. Make your changes
3. Test thoroughly
4. Update documentation if needed
5. Submit a pull request with a clear description

## Project Structure

```
tea-timer/
├── main.go          # Main application code
├── go.mod           # Go module file
├── go.sum           # Dependency checksums
├── README.md        # Project documentation
├── CONTRIBUTING.md  # This file
└── LICENSE          # MIT License
```

## Areas for Contribution

### Features
- Sound notifications
- Multiple timer presets
- Timer history
- Export/import timer configurations
- Additional UI themes

### Improvements
- Performance optimizations
- Better error handling
- Enhanced documentation
- Additional test coverage

### Bug Fixes
- UI rendering issues
- Timer accuracy problems
- Cross-platform compatibility

## Release Process

When contributing to releases:

1. Update version in relevant files
2. Build binaries for all platforms:
   ```bash
   # Linux
   GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o tea-timer-linux-amd64 main.go
   
   # Windows
   GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o tea-timer-windows-amd64.exe main.go
   
   # macOS
   GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o tea-timer-darwin-amd64 main.go
   ```
3. Update `.gitignore` if needed
4. Create git tag for the release

## Questions or Issues?

- Open an issue for bugs or feature requests
- Use discussions for questions and ideas
- Join the community discussions

## License

By contributing to Tea Timer, you agree that your contributions will be licensed under the MIT License. 