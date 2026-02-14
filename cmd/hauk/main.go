package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mnesler/hauk-tui/internal/app"
	"github.com/mnesler/hauk-tui/internal/config"
	"github.com/mnesler/hauk-tui/internal/ui"
)

func main() {
	// Load config and set initial theme
	cfg, err := config.Load()
	if err != nil {
		cfg = config.DefaultConfig()
	}
	ui.SetActiveTheme(cfg.Theme)

	// Create the initial model
	m := app.NewModel()

	// Start the Bubble Tea program
	p := tea.NewProgram(
		m,
		tea.WithAltScreen(),       // Use alternate screen buffer
		tea.WithMouseCellMotion(), // Enable mouse support
	)

	// Run the program
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
