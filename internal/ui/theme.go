package ui

import "github.com/charmbracelet/lipgloss"

// Theme represents a complete color scheme for the application
type Theme struct {
	Name          string
	ChatBg        lipgloss.Color
	DiagramBg     lipgloss.Color
	UserMsgBg     lipgloss.Color
	AgentMsgBg    lipgloss.Color
	InputBg       lipgloss.Color
	TextPrimary   lipgloss.Color
	TextSecondary lipgloss.Color
	TextMuted     lipgloss.Color
	AccentUser    lipgloss.Color
	AccentAgent   lipgloss.Color
	AccentCode    lipgloss.Color
}

// ActiveTheme is the currently active theme
var ActiveTheme = CatppuccinMocha

// Apply applies this theme as the active theme
func (t *Theme) Apply() {
	ActiveTheme = t
}

// Theme Definitions

// CatppuccinMocha - Soft purple-gray tones
var CatppuccinMocha = &Theme{
	Name:          "catppuccin-mocha",
	ChatBg:        lipgloss.Color("#1e1e2e"),
	DiagramBg:     lipgloss.Color("#252538"),
	UserMsgBg:     lipgloss.Color("#2a2a3c"),
	AgentMsgBg:    lipgloss.Color("#313244"),
	InputBg:       lipgloss.Color("#1e1e2e"),
	TextPrimary:   lipgloss.Color("#cdd6f4"),
	TextSecondary: lipgloss.Color("#a6adc8"),
	TextMuted:     lipgloss.Color("#6c7086"),
	AccentUser:    lipgloss.Color("#89b4fa"),
	AccentAgent:   lipgloss.Color("#a6e3a1"),
	AccentCode:    lipgloss.Color("#f5c2e7"),
}

// Dracula - Purple and pink vampire theme
var Dracula = &Theme{
	Name:          "dracula",
	ChatBg:        lipgloss.Color("#282a36"),
	DiagramBg:     lipgloss.Color("#2f3241"),
	UserMsgBg:     lipgloss.Color("#44475a"),
	AgentMsgBg:    lipgloss.Color("#383a4a"),
	InputBg:       lipgloss.Color("#282a36"),
	TextPrimary:   lipgloss.Color("#f8f8f2"),
	TextSecondary: lipgloss.Color("#bfbfbf"),
	TextMuted:     lipgloss.Color("#6272a4"),
	AccentUser:    lipgloss.Color("#8be9fd"),
	AccentAgent:   lipgloss.Color("#50fa7b"),
	AccentCode:    lipgloss.Color("#ff79c6"),
}

// Nord - Arctic blue-tinted theme
var Nord = &Theme{
	Name:          "nord",
	ChatBg:        lipgloss.Color("#2e3440"),
	DiagramBg:     lipgloss.Color("#3b4252"),
	UserMsgBg:     lipgloss.Color("#434c5e"),
	AgentMsgBg:    lipgloss.Color("#4c566a"),
	InputBg:       lipgloss.Color("#2e3440"),
	TextPrimary:   lipgloss.Color("#eceff4"),
	TextSecondary: lipgloss.Color("#d8dee9"),
	TextMuted:     lipgloss.Color("#616e88"),
	AccentUser:    lipgloss.Color("#88c0d0"),
	AccentAgent:   lipgloss.Color("#a3be8c"),
	AccentCode:    lipgloss.Color("#b48ead"),
}

// Gruvbox - Warm retro brown tones
var Gruvbox = &Theme{
	Name:          "gruvbox",
	ChatBg:        lipgloss.Color("#282828"),
	DiagramBg:     lipgloss.Color("#32302f"),
	UserMsgBg:     lipgloss.Color("#3c3836"),
	AgentMsgBg:    lipgloss.Color("#504945"),
	InputBg:       lipgloss.Color("#282828"),
	TextPrimary:   lipgloss.Color("#ebdbb2"),
	TextSecondary: lipgloss.Color("#d5c4a1"),
	TextMuted:     lipgloss.Color("#928374"),
	AccentUser:    lipgloss.Color("#83a598"),
	AccentAgent:   lipgloss.Color("#b8bb26"),
	AccentCode:    lipgloss.Color("#d3869b"),
}

// TokyoNight - Deep blue/purple night theme
var TokyoNight = &Theme{
	Name:          "tokyo-night",
	ChatBg:        lipgloss.Color("#1a1b26"),
	DiagramBg:     lipgloss.Color("#24283b"),
	UserMsgBg:     lipgloss.Color("#292e42"),
	AgentMsgBg:    lipgloss.Color("#343a55"),
	InputBg:       lipgloss.Color("#1a1b26"),
	TextPrimary:   lipgloss.Color("#c0caf5"),
	TextSecondary: lipgloss.Color("#a9b1d6"),
	TextMuted:     lipgloss.Color("#565f89"),
	AccentUser:    lipgloss.Color("#7aa2f7"),
	AccentAgent:   lipgloss.Color("#9ece6a"),
	AccentCode:    lipgloss.Color("#bb9af7"),
}

// GitHubDark - GitHub's dark theme
var GitHubDark = &Theme{
	Name:          "github-dark",
	ChatBg:        lipgloss.Color("#0d1117"),
	DiagramBg:     lipgloss.Color("#161b22"),
	UserMsgBg:     lipgloss.Color("#21262d"),
	AgentMsgBg:    lipgloss.Color("#2d333b"),
	InputBg:       lipgloss.Color("#0d1117"),
	TextPrimary:   lipgloss.Color("#c9d1d9"),
	TextSecondary: lipgloss.Color("#8b949e"),
	TextMuted:     lipgloss.Color("#6e7681"),
	AccentUser:    lipgloss.Color("#58a6ff"),
	AccentAgent:   lipgloss.Color("#3fb950"),
	AccentCode:    lipgloss.Color("#d2a8ff"),
}

// BlueMonochromeDark - Trendy blue monochromatic theme with dark background
var BlueMonochromeDark = &Theme{
	Name:          "blue-monochrome-dark",
	ChatBg:        lipgloss.Color("#0041C7"), // Crayola's Absolute Zero (darkest)
	DiagramBg:     lipgloss.Color("#0160C9"), // True Blue (darker)
	UserMsgBg:     lipgloss.Color("#0D85D8"), // Blue Cola (medium)
	AgentMsgBg:    lipgloss.Color("#0160C9"), // True Blue (darker)
	InputBg:       lipgloss.Color("#0041C7"), // Crayola's Absolute Zero (darkest)
	TextPrimary:   lipgloss.Color("#3ACBE8"), // Picton Blue (brightest)
	TextSecondary: lipgloss.Color("#1CA3DE"), // Battery Charged Blue (bright)
	TextMuted:     lipgloss.Color("#0D85D8"), // Blue Cola (medium)
	AccentUser:    lipgloss.Color("#3ACBE8"), // Picton Blue (brightest)
	AccentAgent:   lipgloss.Color("#1CA3DE"), // Battery Charged Blue (bright)
	AccentCode:    lipgloss.Color("#0D85D8"), // Blue Cola (medium)
}

// BlueMonochrome - Trendy blue monochromatic theme with balanced contrast
var BlueMonochrome = &Theme{
	Name:          "blue-monochrome",
	ChatBg:        lipgloss.Color("#0160C9"), // True Blue (darker)
	DiagramBg:     lipgloss.Color("#0D85D8"), // Blue Cola (medium)
	UserMsgBg:     lipgloss.Color("#1CA3DE"), // Battery Charged Blue (bright)
	AgentMsgBg:    lipgloss.Color("#0D85D8"), // Blue Cola (medium)
	InputBg:       lipgloss.Color("#0160C9"), // True Blue (darker)
	TextPrimary:   lipgloss.Color("#FFFFFF"), // White (maximum readability)
	TextSecondary: lipgloss.Color("#3ACBE8"), // Picton Blue (brightest)
	TextMuted:     lipgloss.Color("#0041C7"), // Crayola's Absolute Zero (darkest)
	AccentUser:    lipgloss.Color("#3ACBE8"), // Picton Blue (brightest)
	AccentAgent:   lipgloss.Color("#1CA3DE"), // Battery Charged Blue (bright)
	AccentCode:    lipgloss.Color("#0D85D8"), // Blue Cola (medium)
}
