package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mnesler/hauk-tui/internal/app"
	"github.com/mnesler/hauk-tui/internal/config"
	"github.com/mnesler/hauk-tui/internal/logger"
	"github.com/mnesler/hauk-tui/internal/ui"
)

func main() {
	// Initialize logger with 1000 entry buffer
	logger.Init(1000)
	logger.StartupMessage("v0.1.0")

	// Load config and set initial theme
	cfg, err := config.Load()
	if err != nil {
		cfg = config.DefaultConfig()
		logger.Component("config").Warn("Failed to load config, using defaults")
	} else {
		logger.Component("config").Infof("Loaded config: theme=%s", cfg.Theme)
	}
	ui.SetActiveTheme(cfg.Theme)

	// Create the initial model
	m := app.NewModel()
	logger.Component("app").Info("Application model created")

	// Start the Bubble Tea program
	p := tea.NewProgram(
		m,
		tea.WithAltScreen(),       // Use alternate screen buffer
		tea.WithMouseCellMotion(), // Enable mouse support
	)

	logger.Component("app").Info("Starting Bubble Tea program...")

	// Run the program
	if _, err := p.Run(); err != nil {
		logger.Component("app").Errorf("Program error: %v", err)
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	logger.Component("app").Info("Application exiting normally")
}
