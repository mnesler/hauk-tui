package app

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mnesler/hauk-tui/internal/chat"
	"github.com/mnesler/hauk-tui/internal/ui"
)

func TestNewModel(t *testing.T) {
	m := NewModel()
	
	if m.messages == nil {
		t.Error("NewModel() messages is nil")
	}
	
	if m.config == nil {
		t.Error("NewModel() config is nil")
	}
	
	if m.showThemeSelector {
		t.Error("NewModel() showThemeSelector should be false")
	}
	
	if m.config.Theme != m.savedTheme {
		t.Errorf("NewModel() savedTheme = %q, config.Theme = %q, should match", m.savedTheme, m.config.Theme)
	}
}

func TestUpdate_WindowSize(t *testing.T) {
	m := NewModel()
	
	// Send window size message
	msg := tea.WindowSizeMsg{Width: 100, Height: 50}
	newModel, _ := m.Update(msg)
	m = newModel.(Model)
	
	if m.width != 100 {
		t.Errorf("After WindowSizeMsg, width = %d, want 100", m.width)
	}
	
	if m.height != 50 {
		t.Errorf("After WindowSizeMsg, height = %d, want 50", m.height)
	}
	
	// Check panel widths calculated correctly (50/50 split)
	if m.chatWidth != 50 {
		t.Errorf("After WindowSizeMsg, chatWidth = %d, want 50", m.chatWidth)
	}
	
	if m.diagramWidth != 50 {
		t.Errorf("After WindowSizeMsg, diagramWidth = %d, want 50", m.diagramWidth)
	}
}

func TestUpdate_ThemeCommand(t *testing.T) {
	m := NewModel()
	m.width = 100
	m.height = 50
	
	// Set input value to /theme
	m.input.SetValue("/theme")
	
	// Send Enter key
	msg := tea.KeyMsg{Type: tea.KeyEnter}
	newModel, _ := m.Update(msg)
	m = newModel.(Model)
	
	// Should show theme selector
	if !m.showThemeSelector {
		t.Error("After /theme command, showThemeSelector should be true")
	}
	
	// Input should be cleared
	if m.input.Value() != "" {
		t.Errorf("After /theme command, input value = %q, want empty", m.input.Value())
	}
	
	// Input should be blurred
	if m.input.Focused() {
		t.Error("After /theme command, input should be blurred")
	}
}

func TestUpdate_RegularMessage(t *testing.T) {
	m := NewModel()
	m.width = 100
	m.height = 50
	
	// Set input value to regular message
	m.input.SetValue("hello world")
	
	// Send Enter key
	msg := tea.KeyMsg{Type: tea.KeyEnter}
	newModel, _ := m.Update(msg)
	m = newModel.(Model)
	
	// Should not show theme selector
	if m.showThemeSelector {
		t.Error("After regular message, showThemeSelector should be false")
	}
	
	// Should add message to history
	if len(m.messages) != 1 {
		t.Errorf("After regular message, messages length = %d, want 1", len(m.messages))
	}
	
	if len(m.messages) > 0 {
		if m.messages[0].Role != chat.RoleUser {
			t.Errorf("Message role = %v, want %v", m.messages[0].Role, chat.RoleUser)
		}
		
		if m.messages[0].Content != "hello world" {
			t.Errorf("Message content = %q, want %q", m.messages[0].Content, "hello world")
		}
	}
	
	// Input should be cleared
	if m.input.Value() != "" {
		t.Errorf("After sending message, input value = %q, want empty", m.input.Value())
	}
}

func TestUpdate_ThemeSelector_Cancel(t *testing.T) {
	m := NewModel()
	m.width = 100
	m.height = 50
	
	// Set up theme selector state
	originalTheme := "catppuccin-mocha"
	m.config.Theme = originalTheme
	m.savedTheme = originalTheme
	m.showThemeSelector = true
	m.input.Blur()
	
	// Change to different theme (simulate preview)
	ui.SetActiveTheme("dracula")
	
	// Send Esc key
	msg := tea.KeyMsg{Type: tea.KeyEsc}
	newModel, _ := m.Update(msg)
	m = newModel.(Model)
	
	// Should hide theme selector
	if m.showThemeSelector {
		t.Error("After Esc, showThemeSelector should be false")
	}
	
	// Should revert to original theme
	if ui.ActiveTheme.Name != originalTheme {
		t.Errorf("After Esc, ActiveTheme.Name = %q, want %q", ui.ActiveTheme.Name, originalTheme)
	}
	
	// Input should be focused again
	if !m.input.Focused() {
		t.Error("After Esc, input should be focused")
	}
	
	// Input should be empty
	if m.input.Value() != "" {
		t.Errorf("After Esc, input value = %q, want empty", m.input.Value())
	}
}

func TestUpdate_ThemeSelector_Select(t *testing.T) {
	// Save original HOME for config
	// (We can't easily test config saving without mocking, but we can test the logic)
	
	m := NewModel()
	m.width = 100
	m.height = 50
	
	// Set up theme selector state
	m = m.showThemeSelectorModal()
	
	// Verify selector is shown
	if !m.showThemeSelector {
		t.Fatal("showThemeSelectorModal() did not set showThemeSelector")
	}
	
	// The list should have items
	if m.themeList.Items() == nil || len(m.themeList.Items()) == 0 {
		t.Fatal("Theme list has no items")
	}
}

func TestFormatThemeName(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"catppuccin-mocha", "Catppuccin Mocha"},
		{"dracula", "Dracula"},
		{"nord", "Nord"},
		{"gruvbox", "Gruvbox"},
		{"tokyo-night", "Tokyo Night"},
		{"github-dark", "GitHub Dark"},
		{"unknown", "unknown"}, // Fallback
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := formatThemeName(tt.input)
			if got != tt.want {
				t.Errorf("formatThemeName(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestShowThemeSelectorModal(t *testing.T) {
	m := NewModel()
	m.width = 100
	m.height = 50
	
	originalTheme := m.config.Theme
	
	m = m.showThemeSelectorModal()
	
	// Should show selector
	if !m.showThemeSelector {
		t.Error("showThemeSelectorModal() did not set showThemeSelector")
	}
	
	// Should save current theme
	if m.savedTheme != originalTheme {
		t.Errorf("savedTheme = %q, want %q", m.savedTheme, originalTheme)
	}
	
	// Should blur input
	if m.input.Focused() {
		t.Error("showThemeSelectorModal() did not blur input")
	}
	
	// Should populate list with themes
	items := m.themeList.Items()
	if len(items) != 6 {
		t.Errorf("Theme list has %d items, want 6", len(items))
	}
	
	// Verify all items are valid themeItems
	for i, item := range items {
		if themeItem, ok := item.(themeItem); ok {
			if themeItem.name == "" {
				t.Errorf("Item %d has empty name", i)
			}
			if themeItem.displayName == "" {
				t.Errorf("Item %d has empty displayName", i)
			}
		} else {
			t.Errorf("Item %d is not a themeItem", i)
		}
	}
}

func TestUpdate_EmptyInput(t *testing.T) {
	m := NewModel()
	m.width = 100
	m.height = 50
	
	// Set empty input
	m.input.SetValue("")
	
	// Send Enter key
	msg := tea.KeyMsg{Type: tea.KeyEnter}
	newModel, _ := m.Update(msg)
	m = newModel.(Model)
	
	// Should not add message
	if len(m.messages) != 0 {
		t.Errorf("After empty input, messages length = %d, want 0", len(m.messages))
	}
	
	// Should not show theme selector
	if m.showThemeSelector {
		t.Error("After empty input, showThemeSelector should be false")
	}
}

func TestUpdate_AltEnter(t *testing.T) {
	m := NewModel()
	m.width = 100
	m.height = 50
	
	// Set input value
	initialValue := "line1"
	m.input.SetValue(initialValue)
	
	// Send Alt+Enter
	msg := tea.KeyMsg{Type: tea.KeyEnter, Alt: true}
	newModel, _ := m.Update(msg)
	m = newModel.(Model)
	
	// Should add newline to input (or space, depending on textinput implementation)
	// The key point is it should NOT send the message
	if len(m.messages) != 0 {
		t.Errorf("After Alt+Enter, messages length = %d, want 0 (message should not be sent)", len(m.messages))
	}
	
	// Input value should be modified (either with newline or space added)
	if m.input.Value() == initialValue {
		t.Error("After Alt+Enter, input value should be modified")
	}
}

func TestThemeItem_Interface(t *testing.T) {
	// Test that themeItem implements list.Item interface
	item := themeItem{
		name:        "test",
		displayName: "Test Theme",
	}
	
	if item.FilterValue() != "Test Theme" {
		t.Errorf("FilterValue() = %q, want %q", item.FilterValue(), "Test Theme")
	}
	
	if item.Title() != "Test Theme" {
		t.Errorf("Title() = %q, want %q", item.Title(), "Test Theme")
	}
	
	if item.Description() != "" {
		t.Errorf("Description() = %q, want empty", item.Description())
	}
}

// Benchmark update operations
func BenchmarkUpdate_WindowSize(b *testing.B) {
	m := NewModel()
	msg := tea.WindowSizeMsg{Width: 100, Height: 50}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Update(msg)
	}
}

func BenchmarkUpdate_ThemeCommand(b *testing.B) {
	m := NewModel()
	m.width = 100
	m.height = 50
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.input.SetValue("/theme")
		msg := tea.KeyMsg{Type: tea.KeyEnter}
		m.Update(msg)
		
		// Reset state
		m.showThemeSelector = false
		m.input.Focus()
	}
}
