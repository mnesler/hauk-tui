package theme

import "github.com/charmbracelet/lipgloss"

// Theme represents a complete color scheme and styling for the application
type Theme struct {
	Name string
	Colors Colors
	Styles Styles
}

// Colors defines all the color values used in the theme
type Colors struct {
	// Background colors
	ChatBg    lipgloss.Color
	DiagramBg lipgloss.Color
	
	// Message colors
	UserMsgBg  lipgloss.Color
	AgentMsgBg lipgloss.Color
	InputBg    lipgloss.Color
	
	// Text colors
	TextPrimary   lipgloss.Color
	TextSecondary lipgloss.Color
	TextMuted     lipgloss.Color
	
	// Accent colors
	AccentUser  lipgloss.Color
	AccentAgent lipgloss.Color
	AccentCode  lipgloss.Color
}

// Styles defines all the lipgloss styles used in the application
type Styles struct {
	UserMsgStyle      lipgloss.Style
	AgentMsgStyle     lipgloss.Style
	CodeStyle         lipgloss.Style
	ChatPanelStyle    lipgloss.Style
	DiagramPanelStyle lipgloss.Style
	InputStyle        lipgloss.Style
}