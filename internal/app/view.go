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

	// If theme selector is shown, render it over the main view
	if m.showThemeSelector {
		// Render theme selector modal
		themeSelector := m.renderThemeSelector()
		
		// Overlay the selector centered on screen
		// The background will be the empty space around it
		overlay := lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			themeSelector,
		)
		
		return overlay
	}

	return m.renderMainView()
}

// renderMainView renders the main application view
func (m Model) renderMainView() string {
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
	header := ui.GetHeaderStyle(ui.ActiveTheme.ChatBg).
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
	return ui.GetChatPanelStyle(m.chatWidth, m.height-3).
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
		style = ui.GetUserMsgStyle(m.chatWidth - 4)
		prefix = fmt.Sprintf("[%s] You:", timestamp)
	case chat.RoleAgent:
		style = ui.GetAgentMsgStyle(m.chatWidth - 4)
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
	header := ui.GetHeaderStyle(ui.ActiveTheme.DiagramBg).
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
	diagramStyled := ui.GetTextSecondaryStyle().
		Padding(1, 2).
		Render(diagramContent)

	content = append(content, diagramStyled)

	// Join content
	panelContent := lipgloss.JoinVertical(
		lipgloss.Left,
		content...,
	)

	// Apply panel styling
	return ui.GetDiagramPanelStyle(m.diagramWidth, m.height-3).
		Render(panelContent)
}

// renderInputBar renders the input bar at the bottom
func (m Model) renderInputBar() string {
	inputView := m.input.View()

	return ui.GetInputStyle(m.width).
		Render("> " + inputView)
}

// renderThemeSelector renders the theme selector modal
func (m Model) renderThemeSelector() string {
	// Modal dimensions
	modalWidth := 50
	modalHeight := 20

	// Instructions
	instructions := ui.GetTextMutedStyle().
		Render("↑/↓: navigate • Enter: select • Esc: cancel")

	// List view
	listView := m.themeList.View()

	// Combine instructions and list
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		listView,
		"",
		instructions,
	)

	// Modal style
	modalStyle := lipgloss.NewStyle().
		Background(ui.ActiveTheme.ChatBg).
		Foreground(ui.ActiveTheme.TextPrimary).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ui.ActiveTheme.AccentUser).
		Padding(1, 2).
		Width(modalWidth).
		Height(modalHeight)

	return modalStyle.Render(content)
}

// Helper to check if string contains substring
func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}
