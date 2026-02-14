package ui

import "github.com/charmbracelet/lipgloss"

// Color scheme - Catppuccin Mocha (soft, similar tones)
var (
	// Background colors (subtle distinction between panes)
	ChatBg    = lipgloss.Color("#1e1e2e") // Deep purple-gray
	DiagramBg = lipgloss.Color("#252538") // Slightly lighter purple

	// Message colors
	UserMsgBg  = lipgloss.Color("#2a2a3c") // Subtle highlight
	AgentMsgBg = lipgloss.Color("#313244") // Different subtle highlight
	InputBg    = lipgloss.Color("#1e1e2e") // Match chat panel

	// Text colors
	TextPrimary   = lipgloss.Color("#cdd6f4") // Soft white
	TextSecondary = lipgloss.Color("#a6adc8") // Muted gray
	TextMuted     = lipgloss.Color("#6c7086") // Placeholder text

	// Accent colors (minimal usage)
	AccentUser  = lipgloss.Color("#89b4fa") // Blue for user
	AccentAgent = lipgloss.Color("#a6e3a1") // Green for agent
	AccentCode  = lipgloss.Color("#f5c2e7") // Pink for code elements
)

// Styles
var (
	// User message style
	UserMsgStyle = lipgloss.NewStyle().
		Background(UserMsgBg).
		Foreground(TextPrimary).
		Padding(1, 2).
		MarginBottom(1)

	// Agent message style
	AgentMsgStyle = lipgloss.NewStyle().
		Background(AgentMsgBg).
		Foreground(TextPrimary).
		Padding(1, 2).
		MarginBottom(1)

	// Code block style
	CodeStyle = lipgloss.NewStyle().
		Foreground(AccentCode).
		MarginLeft(2)

	// Chat panel style
	ChatPanelStyle = lipgloss.NewStyle().
		Background(ChatBg).
		Padding(1)

	// Diagram panel style
	DiagramPanelStyle = lipgloss.NewStyle().
		Background(DiagramBg).
		Padding(1)

	// Input style
	InputStyle = lipgloss.NewStyle().
		Background(InputBg).
		Foreground(TextPrimary).
		Padding(1, 2)
)
