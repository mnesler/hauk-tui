package app

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mnesler/hauk-tui/internal/ui"
)

// themeDelegate is a custom list delegate that renders items with theme colors
type themeDelegate struct{}

// newThemeDelegate creates a new theme-aware list delegate
func newThemeDelegate() list.ItemDelegate {
	return themeDelegate{}
}

// Height returns the height of each list item
func (d themeDelegate) Height() int {
	return 1
}

// Spacing returns the spacing between list items
func (d themeDelegate) Spacing() int {
	return 0
}

// Update handles the delegate's update logic
func (d themeDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

// Render renders a list item with theme-aware styling
func (d themeDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	if item == nil {
		return
	}

	// Cast to themeItem
	themeItem, ok := item.(themeItem)
	if !ok {
		return
	}

	// Determine if this item is selected
	isSelected := index == m.Index()

	var str string
	if isSelected {
		// Selected item: bullet indicator + theme name
		// Use AccentUser color on UserMsgBg background
		selectedStyle := lipgloss.NewStyle().
			Foreground(ui.ActiveTheme.AccentUser).
			Background(ui.ActiveTheme.UserMsgBg).
			Bold(true).
			Padding(0, 1)

		str = selectedStyle.Render(fmt.Sprintf("• %s", themeItem.displayName))
	} else {
		// Unselected item: indented theme name
		// Use TextSecondary color (bright blue)
		unselectedStyle := lipgloss.NewStyle().
			Foreground(ui.ActiveTheme.TextSecondary).
			Padding(0, 1)

		str = unselectedStyle.Render(fmt.Sprintf("  %s", themeItem.displayName))
	}

	fmt.Fprint(w, str)
}
