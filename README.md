# Tea Timer

A minimalist terminal-based countdown timer built with Go and Bubble Tea. Perfect for timing tea steeping, meditation sessions, or any quick countdown needs.

## Features

- **Clean Interface**: Beautiful terminal UI with colored text and smooth animations
- **Visual Feedback**: Spinning animation during countdown, flashing alert when time expires
- **Simple Controls**: Start, pause, reset, and quit with single key presses
- **Flexible Duration**: Set custom countdown time via command line argument
- **Cross-Platform**: Works on Windows, macOS, and Linux

## Usage

```bash
# Run with default 60-second timer
go run main.go

# Run with custom duration (in seconds)
go run main.go 120

# Build and run the executable
go build
./tea-timer 300
```

## Controls

- `s` or `space` - Start/pause timer
- `r` - Reset timer to initial duration
- `q`, `ctrl+c`, or `esc` - Quit the application

When the timer finishes, the display will flash to alert you. Press `s` or `space` to restart the timer.

## Requirements

- Go 1.24 or later
- Terminal with color support

## Installation

1. Clone the repository
2. Run `go mod tidy` to install dependencies
3. Build with `go build` or run directly with `go run main.go`

## Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Terminal UI framework
- [Bubbles](https://github.com/charmbracelet/bubbles) - UI components
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Terminal styling 
