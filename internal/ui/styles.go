package ui

import "github.com/charmbracelet/lipgloss"

// GetUserMsgStyle returns the style for user messages
func GetUserMsgStyle(width int) lipgloss.Style {
	return lipgloss.NewStyle().
		Background(ActiveTheme.ChatBg).
		Foreground(ActiveTheme.TextPrimary).
		Padding(1, 2).
		MarginBottom(1).
		Width(width)
}

// GetAgentMsgStyle returns the style for agent messages
func GetAgentMsgStyle(width int) lipgloss.Style {
	return lipgloss.NewStyle().
		Background(ActiveTheme.ChatBg).
		Foreground(ActiveTheme.TextPrimary).
		Padding(1, 2).
		MarginBottom(1).
		Width(width)
}

// GetCodeStyle returns the style for code blocks
func GetCodeStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(ActiveTheme.AccentCode).
		MarginLeft(2)
}

// GetChatPanelStyle returns the style for the chat panel
func GetChatPanelStyle(width, height int) lipgloss.Style {
	return lipgloss.NewStyle().
		Background(ActiveTheme.ChatBg).
		Padding(1).
		Width(width).
		Height(height)
}

// GetDiagramPanelStyle returns the style for the diagram panel
func GetDiagramPanelStyle(width, height int) lipgloss.Style {
	return lipgloss.NewStyle().
		Background(ActiveTheme.DiagramBg).
		Padding(1).
		Width(width).
		Height(height)
}

// GetInputStyle returns the style for the input bar
func GetInputStyle(width int) lipgloss.Style {
	return lipgloss.NewStyle().
		Background(ActiveTheme.InputBg).
		Foreground(ActiveTheme.TextPrimary).
		Padding(1, 2).
		Width(width)
}

// GetHeaderStyle returns the style for panel headers
func GetHeaderStyle(bg lipgloss.Color) lipgloss.Style {
	return lipgloss.NewStyle().
		Background(bg).
		Foreground(ActiveTheme.TextPrimary).
		Bold(true).
		Padding(0, 2)
}

// GetTextSecondaryStyle returns the style for secondary text
func GetTextSecondaryStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(ActiveTheme.TextSecondary)
}

// GetTextMutedStyle returns the style for muted text
func GetTextMutedStyle(bg lipgloss.Color) lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(ActiveTheme.TextMuted).
		Background(bg)
}
