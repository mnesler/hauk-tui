package app

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mnesler/hauk-tui/internal/chat"
)

// Model represents the main application state
type Model struct {
	// UI Components
	chatViewport viewport.Model
	input        textinput.Model

	// State
	messages       []chat.Message
	currentDiagram string
	isLoading      bool
	width          int
	height         int

	// Layout calculations
	chatWidth    int
	diagramWidth int
}

// NewModel creates a new application model
func NewModel() Model {
	// Initialize input
	input := textinput.New()
	input.Placeholder = "Type a message or paste code..."
	input.Focus()

	// Initialize chat viewport
	vp := viewport.New(0, 0)
	vp.SetContent("")

	return Model{
		chatViewport: vp,
		input:        input,
		messages:     make([]chat.Message, 0),
	}
}

// Init initializes the application
func (m Model) Init() tea.Cmd {
	return textinput.Blink
}
