package app

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mnesler/hauk-tui/internal/chat"
	"github.com/mnesler/hauk-tui/internal/command"
	"github.com/mnesler/hauk-tui/internal/config"
	"github.com/mnesler/hauk-tui/internal/ui"
)

// Message types
type (
	// WindowResizeMsg is sent when the window is resized
	WindowResizeMsg struct {
		Width  int
		Height int
	}

	// SendMessageMsg is sent when user submits a message
	SendMessageMsg struct {
		Content string
	}

	// AgentResponseMsg is sent when agent responds
	AgentResponseMsg struct {
		Content string
		Diagram string
	}
)

// Update handles all messages and updates the model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	// If theme selector is active, handle its input first
	if m.showThemeSelector {
		return m.updateThemeSelector(msg)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyEnter:
			// Check if Alt is pressed
			if msg.Alt {
				// Alt+Enter: add newline
				m.input.SetValue(m.input.Value() + "\n")
			} else {
				// Enter: send message
				content := m.input.Value()
				if content != "" {
					// Check if it's a slash command
					cmdType, _ := command.ParseCommand(content)
					if cmdType == command.CommandTheme {
						// Show theme selector
						m = m.showThemeSelectorModal()
						return m, nil
					} else if cmdType == command.CommandNone {
						// Add user message
						m.messages = append(m.messages, chat.NewMessage(chat.RoleUser, content))
						m.input.SetValue("")

						// Simulate agent response (will be replaced with real LLM call)
						cmds = append(cmds, m.simulateAgentResponse())
					}
				}
			}
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		// Calculate panel widths (50/50 split)
		m.chatWidth = m.width / 2
		m.diagramWidth = m.width - m.chatWidth

		// Update viewport size
		m.chatViewport.Width = m.chatWidth - 2
		m.chatViewport.Height = m.height - 4

		// Update theme list size
		m.themeList.SetSize(40, 12)

	case AgentResponseMsg:
		// Add agent message
		agentMsg := chat.NewMessage(chat.RoleAgent, msg.Content)
		agentMsg.Diagram = msg.Diagram
		m.messages = append(m.messages, agentMsg)

		// Update current diagram if provided
		if msg.Diagram != "" {
			m.currentDiagram = msg.Diagram
		}
	}

	// Update input
	newInput, inputCmd := m.input.Update(msg)
	m.input = newInput
	cmds = append(cmds, inputCmd)

	// Update viewport
	newVp, vpCmd := m.chatViewport.Update(msg)
	m.chatViewport = newVp
	cmds = append(cmds, vpCmd)

	return m, tea.Batch(cmds...)
}

// simulateAgentResponse simulates an agent response (placeholder)
func (m Model) simulateAgentResponse() tea.Cmd {
	return func() tea.Msg {
		// This will be replaced with real LLM integration
		return AgentResponseMsg{
			Content: "I'll create a flowchart for you. Here's a simple example:\n\n```mermaid\ngraph TD\n    A[Start] --> B{Is it working?}\n    B -->|Yes| C[Great!]\n    B -->|No| D[Debug]\n    D --> B\n```",
			Diagram: "graph TD\n    A[Start] --> B{Is it working?}\n    B -->|Yes| C[Great!]\n    B -->|No| D[Debug]\n    D --> B",
		}
	}
}

// showThemeSelectorModal shows the theme selector
func (m Model) showThemeSelectorModal() Model {
	// Get all available themes
	themes := ui.GetAvailableThemes()

	// Create list items
	items := make([]list.Item, len(themes))
	for i, name := range themes {
		theme := ui.GetTheme(name)
		items[i] = themeItem{
			name:        name,
			displayName: theme.Name,
		}
	}

	// Update theme list
	m.themeList.SetItems(items)
	m.themeList.SetSize(40, 12)

	// Save current theme and show selector
	m.savedTheme = m.config.Theme
	m.showThemeSelector = true
	m.input.Blur()

	return m
}

// updateThemeSelector handles input when theme selector is active
func (m Model) updateThemeSelector(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			// Cancel and revert to saved theme
			ui.SetActiveTheme(m.savedTheme)
			m.showThemeSelector = false
			m.input.Focus()
			m.input.SetValue("")
			return m, nil

		case tea.KeyEnter:
			// Apply selected theme
			if item, ok := m.themeList.SelectedItem().(themeItem); ok {
				m.config.Theme = item.name
				ui.SetActiveTheme(item.name)
				
				// Save config
				if err := config.Save(m.config); err != nil {
					// Handle error (could add error message to UI)
				}
			}
			m.showThemeSelector = false
			m.input.Focus()
			m.input.SetValue("")
			return m, nil

		case tea.KeyUp, tea.KeyDown:
			// Update list and apply live preview
			var cmd tea.Cmd
			m.themeList, cmd = m.themeList.Update(msg)
			
			// Apply theme preview
			if item, ok := m.themeList.SelectedItem().(themeItem); ok {
				ui.SetActiveTheme(item.name)
			}
			return m, cmd
		}
	}

	// Update list for other keys
	var cmd tea.Cmd
	m.themeList, cmd = m.themeList.Update(msg)
	return m, cmd
}
