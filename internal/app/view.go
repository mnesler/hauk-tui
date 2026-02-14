package app

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mnesler/hauk-tui/internal/chat"
	"github.com/mnesler/hauk-tui/internal/diagram"
	"github.com/mnesler/hauk-tui/internal/ui"
)

// View renders the UI
func (m Model) View() string {
	if m.width == 0 || m.height == 0 {
		return "Loading..."
	}

	// Render chat panel (left 50%)
	chatPanel := m.renderChatPanel()

	// Render diagram panel (right 50%)
	diagramPanel := m.renderDiagramPanel()

	// Join panels horizontally
	content := lipgloss.JoinHorizontal(
		lipgloss.Top,
		chatPanel,
		diagramPanel,
	)

	// Add input bar at bottom
	inputBar := m.renderInputBar()

	// Combine content and input
	return lipgloss.JoinVertical(
		lipgloss.Left,
		content,
		inputBar,
	)
}

// renderChatPanel renders the left panel with chat messages
func (m Model) renderChatPanel() string {
	var messages []string

	// Header
	header := lipgloss.NewStyle().
		Background(ui.ChatBg).
		Foreground(ui.TextPrimary).
		Bold(true).
		Padding(0, 2).
		Render("Chat")
	messages = append(messages, header)

	// Messages
	for _, msg := range m.messages {
		rendered := m.renderMessage(msg)
		messages = append(messages, rendered)
	}

	// Join all messages
	chatContent := lipgloss.JoinVertical(
		lipgloss.Left,
		messages...,
	)

	// Apply panel styling
	return ui.ChatPanelStyle.
		Width(m.chatWidth).
		Height(m.height - 3).
		Render(chatContent)
}

// renderMessage renders a single chat message
func (m Model) renderMessage(msg chat.Message) string {
	// Format timestamp
	timestamp := msg.Timestamp.Format("15:04")

	// Choose style based on role
	var style lipgloss.Style
	var prefix string

	switch msg.Role {
	case chat.RoleUser:
		style = ui.UserMsgStyle.Width(m.chatWidth - 4)
		prefix = fmt.Sprintf("[%s] You:", timestamp)
	case chat.RoleAgent:
		style = ui.AgentMsgStyle.Width(m.chatWidth - 4)
		prefix = fmt.Sprintf("[%s] Agent:", timestamp)
	}

	// Render content
	content := fmt.Sprintf("%s\n%s", prefix, msg.Content)
	return style.Render(content)
}

// renderDiagramPanel renders the right panel with ASCII diagram
func (m Model) renderDiagramPanel() string {
	var content []string

	// Header
	header := lipgloss.NewStyle().
		Background(ui.DiagramBg).
		Foreground(ui.TextPrimary).
		Bold(true).
		Padding(0, 2).
		Render("Diagram Preview")
	content = append(content, header)

	// Diagram or placeholder
	var diagramContent string
	if m.currentDiagram != "" {
		// Render mermaid to ASCII
		ascii, err := diagram.Render(m.currentDiagram)
		if err != nil {
			diagramContent = fmt.Sprintf("Error rendering diagram: %v", err)
		} else {
			diagramContent = ascii
		}
	} else {
		diagramContent = "No diagram yet. Chat with the agent to generate one!"
	}

	// Style the diagram content
	diagramStyled := lipgloss.NewStyle().
		Foreground(ui.TextSecondary).
		Padding(1, 2).
		Render(diagramContent)

	content = append(content, diagramStyled)

	// Join content
	panelContent := lipgloss.JoinVertical(
		lipgloss.Left,
		content...,
	)

	// Apply panel styling
	return ui.DiagramPanelStyle.
		Width(m.diagramWidth).
		Height(m.height - 3).
		Render(panelContent)
}

// renderInputBar renders the input bar at the bottom
func (m Model) renderInputBar() string {
	inputView := m.input.View()

	return ui.InputStyle.
		Width(m.width).
		Render("> " + inputView)
}

// Helper to check if string contains substring
func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}
