package app

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mnesler/hauk-tui/internal/chat"
	"github.com/mnesler/hauk-tui/internal/config"
)

// themeItem represents a theme in the list
type themeItem struct {
	name        string
	displayName string
}

func (t themeItem) FilterValue() string { return t.displayName }
func (t themeItem) Title() string       { return t.displayName }
func (t themeItem) Description() string { return "" }

// Model represents the main application state
type Model struct {
	// UI Components
	chatViewport viewport.Model
	input        textinput.Model
	themeList    list.Model

	// State
	messages       []chat.Message
	currentDiagram string
	isLoading      bool
	width          int
	height         int

	// Theme state
	config            *config.Config
	showThemeSelector bool
	previewTheme      string
	savedTheme        string

	// Layout calculations
	chatWidth    int
	diagramWidth int
}

// NewModel creates a new application model
func NewModel() Model {
	// Load config
	cfg, err := config.Load()
	if err != nil {
		cfg = config.DefaultConfig()
	}

	// Initialize input
	input := textinput.New()
	input.Placeholder = "Type a message or paste code..."
	input.Focus()

	// Initialize chat viewport
	vp := viewport.New(0, 0)
	vp.SetContent("")

	// Initialize theme list (will be populated when shown)
	themeList := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	themeList.Title = "Select Theme"

	return Model{
		chatViewport:      vp,
		input:             input,
		themeList:         themeList,
		messages:          make([]chat.Message, 0),
		config:            cfg,
		showThemeSelector: false,
		previewTheme:      cfg.Theme,
		savedTheme:        cfg.Theme,
	}
}

// Init initializes the application
func (m Model) Init() tea.Cmd {
	return textinput.Blink
}
