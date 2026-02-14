package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mnesler/hauk-tui/internal/chat"
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

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyEnter:
			// Check if Shift is pressed using modifiers
			if msg.Alt {
				// Alt+Enter: add newline
				m.input.SetValue(m.input.Value() + "\n")
			} else {
				// Enter: send message
				content := m.input.Value()
				if content != "" {
					// Add user message
					m.messages = append(m.messages, chat.NewMessage(chat.RoleUser, content))
					m.input.SetValue("")

					// Simulate agent response (will be replaced with real LLM call)
					cmds = append(cmds, m.simulateAgentResponse())
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
