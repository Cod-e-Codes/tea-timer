package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	// Styling
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("69")).
			Bold(true).
			Render

	timeStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252")).
			Bold(true).
			Render

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			Render

	spinnerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("69"))

	statusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("154")).
			Bold(true).
			Render

	flashStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("196")).
			Bold(true).
			Render
)

// tickMsg is sent every second.
type tickMsg time.Time

// flashMsg is sent every 500ms for flashing effect
type flashMsg time.Time

// model holds the state of the countdown timer.
type model struct {
	initial   time.Duration
	remaining time.Duration
	running   bool
	quitting  bool
	spinner   spinner.Model
	finished  bool
	flash     bool
}

// NewModel creates a new countdown model with the given duration.
func NewModel(d time.Duration) model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = spinnerStyle

	return model{
		initial:   d,
		remaining: d,
		running:   false,
		quitting:  false,
		spinner:   s,
		finished:  false,
		flash:     false,
	}
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func flash() tea.Cmd {
	return tea.Tick(500*time.Millisecond, func(t time.Time) tea.Msg {
		return flashMsg(t)
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "s", " ":
			if m.finished {
				// Reset if finished
				m.remaining = m.initial
				m.finished = false
				m.running = false
				return m, m.spinner.Tick
			}
			m.running = !m.running
			if m.running {
				return m, tea.Batch(tick(), m.spinner.Tick)
			}
			return m, m.spinner.Tick
		case "r":
			m.remaining = m.initial
			m.running = false
			m.finished = false
			return m, m.spinner.Tick
		case "q", "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit
		}

	case tickMsg:
		if m.running {
			if m.remaining > 0 {
				m.remaining -= time.Second
				if m.remaining == 0 {
					m.running = false
					m.finished = true
					return m, tea.Batch(flash(), m.spinner.Tick)
				}
				return m, tea.Batch(tick(), m.spinner.Tick)
			}
		}
		return m, m.spinner.Tick

	case flashMsg:
		if m.finished {
			m.flash = !m.flash
			return m, flash()
		}
		return m, nil

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	if m.quitting {
		return ""
	}

	// Format remaining as MM:SS
	min := int(m.remaining.Minutes())
	sec := int(m.remaining.Seconds()) % 60
	timeStr := fmt.Sprintf("%02d:%02d", min, sec)

	// Create the display
	var s string

	// Title
	s += titleStyle("🍵 Tea Timer") + "\n\n"

	// Time display with spinner
	if m.finished {
		// Flash the time when finished
		if m.flash {
			s += fmt.Sprintf(" %s %s\n", "⏰", flashStyle(timeStr))
		} else {
			s += fmt.Sprintf(" %s %s\n", "⏰", timeStyle(timeStr))
		}
		s += flashStyle("⏰  Time's up!") + "\n"
	} else if m.running {
		s += fmt.Sprintf(" %s %s\n", m.spinner.View(), timeStyle(timeStr))
		s += statusStyle("▶️  Running...") + "\n"
	} else {
		s += fmt.Sprintf(" %s %s\n", "⏸️ ", timeStyle(timeStr))
		s += statusStyle("⏸️  Paused") + "\n"
	}

	// Help text
	if m.finished {
		s += "\n" + helpStyle("(s/space) restart • (r) reset • (q) quit")
	} else {
		s += "\n" + helpStyle("(s/space) start/pause • (r) reset • (q) quit")
	}

	return s
}

func main() {
	// Default to 60 seconds or parse first arg as seconds.
	dur := 60 * time.Second
	if len(os.Args) > 1 {
		if secs, err := strconv.Atoi(os.Args[1]); err == nil {
			dur = time.Duration(secs) * time.Second
		}
	}

	p := tea.NewProgram(NewModel(dur))
	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
