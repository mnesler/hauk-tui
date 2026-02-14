package ui

import "sort"

// AvailableThemes maps theme names to their definitions
var AvailableThemes = map[string]*Theme{
	"catppuccin-mocha": CatppuccinMocha,
	"dracula":          Dracula,
	"nord":             Nord,
	"gruvbox":          Gruvbox,
	"tokyo-night":      TokyoNight,
	"github-dark":      GitHubDark,
}

// GetTheme returns a theme by name, or nil if not found
func GetTheme(name string) *Theme {
	return AvailableThemes[name]
}

// GetThemeNames returns a sorted list of all available theme names
func GetThemeNames() []string {
	names := make([]string, 0, len(AvailableThemes))
	for name := range AvailableThemes {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// GetAvailableThemes returns a sorted list of theme names (alias for GetThemeNames)
func GetAvailableThemes() []string {
	return GetThemeNames()
}

// SetActiveTheme sets the active theme by name
func SetActiveTheme(name string) bool {
	theme := GetTheme(name)
	if theme == nil {
		return false
	}
	theme.Apply()
	return true
}
