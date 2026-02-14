package ui

import (
	"testing"
)

func TestGetTheme_Valid(t *testing.T) {
	tests := []struct {
		name      string
		themeName string
		wantNil   bool
	}{
		{"catppuccin-mocha", "catppuccin-mocha", false},
		{"dracula", "dracula", false},
		{"nord", "nord", false},
		{"gruvbox", "gruvbox", false},
		{"tokyo-night", "tokyo-night", false},
		{"github-dark", "github-dark", false},
		{"blue-monochrome-dark", "blue-monochrome-dark", false},
		{"blue-monochrome", "blue-monochrome", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			theme := GetTheme(tt.themeName)

			if tt.wantNil && theme != nil {
				t.Errorf("GetTheme(%q) = %v, want nil", tt.themeName, theme)
			}

			if !tt.wantNil && theme == nil {
				t.Errorf("GetTheme(%q) = nil, want non-nil", tt.themeName)
			}

			if theme != nil && theme.Name != tt.themeName {
				t.Errorf("GetTheme(%q).Name = %q, want %q", tt.themeName, theme.Name, tt.themeName)
			}
		})
	}
}

func TestGetTheme_Invalid(t *testing.T) {
	tests := []string{
		"invalid",
		"nonexistent",
		"",
		"DRACULA", // Case sensitive
		"catppuccin", // Partial name
	}

	for _, themeName := range tests {
		t.Run(themeName, func(t *testing.T) {
			theme := GetTheme(themeName)
			if theme != nil {
				t.Errorf("GetTheme(%q) = %v, want nil for invalid theme", themeName, theme)
			}
		})
	}
}

func TestGetAvailableThemes(t *testing.T) {
	themes := GetAvailableThemes()

	if len(themes) != 8 {
		t.Errorf("GetAvailableThemes() returned %d themes, want 8", len(themes))
	}

	// Check all expected themes are present
	expected := map[string]bool{
		"catppuccin-mocha":     false,
		"dracula":              false,
		"nord":                 false,
		"gruvbox":              false,
		"tokyo-night":          false,
		"github-dark":          false,
		"blue-monochrome-dark": false,
		"blue-monochrome":      false,
	}

	for _, name := range themes {
		if _, ok := expected[name]; ok {
			expected[name] = true
		} else {
			t.Errorf("Unexpected theme in list: %q", name)
		}
	}

	// Check all expected themes were found
	for name, found := range expected {
		if !found {
			t.Errorf("Expected theme %q not found in list", name)
		}
	}
}

func TestGetAvailableThemes_Sorted(t *testing.T) {
	themes := GetAvailableThemes()

	// Check if sorted
	for i := 1; i < len(themes); i++ {
		if themes[i-1] > themes[i] {
			t.Errorf("GetAvailableThemes() not sorted: %q > %q", themes[i-1], themes[i])
		}
	}
}

func TestSetActiveTheme_Valid(t *testing.T) {
	// Save original theme
	originalTheme := ActiveTheme
	defer func() { ActiveTheme = originalTheme }()

	tests := []string{
		"catppuccin-mocha",
		"dracula",
		"nord",
		"gruvbox",
		"tokyo-night",
		"github-dark",
		"blue-monochrome-dark",
		"blue-monochrome",
	}

	for _, themeName := range tests {
		t.Run(themeName, func(t *testing.T) {
			ok := SetActiveTheme(themeName)

			if !ok {
				t.Errorf("SetActiveTheme(%q) = false, want true", themeName)
			}

			if ActiveTheme == nil {
				t.Fatal("ActiveTheme is nil after SetActiveTheme")
			}

			if ActiveTheme.Name != themeName {
				t.Errorf("ActiveTheme.Name = %q, want %q", ActiveTheme.Name, themeName)
			}
		})
	}
}

func TestSetActiveTheme_Invalid(t *testing.T) {
	// Save original theme
	originalTheme := ActiveTheme
	defer func() { ActiveTheme = originalTheme }()

	tests := []string{
		"invalid",
		"nonexistent",
		"",
	}

	for _, themeName := range tests {
		t.Run(themeName, func(t *testing.T) {
			ok := SetActiveTheme(themeName)

			if ok {
				t.Errorf("SetActiveTheme(%q) = true, want false for invalid theme", themeName)
			}

			// ActiveTheme should remain unchanged
			if ActiveTheme != originalTheme {
				t.Error("ActiveTheme changed after invalid SetActiveTheme")
			}
		})
	}
}

func TestTheme_Apply(t *testing.T) {
	// Save original theme
	originalTheme := ActiveTheme
	defer func() { ActiveTheme = originalTheme }()

	theme := GetTheme("dracula")
	if theme == nil {
		t.Fatal("Could not get dracula theme")
	}

	theme.Apply()

	if ActiveTheme != theme {
		t.Error("Theme.Apply() did not set ActiveTheme")
	}

	if ActiveTheme.Name != "dracula" {
		t.Errorf("After Apply(), ActiveTheme.Name = %q, want %q", ActiveTheme.Name, "dracula")
	}
}

func TestTheme_HasAllColors(t *testing.T) {
	// Test that all themes have all required color fields set
	themes := GetAvailableThemes()

	for _, name := range themes {
		t.Run(name, func(t *testing.T) {
			theme := GetTheme(name)
			if theme == nil {
				t.Fatalf("GetTheme(%q) returned nil", name)
			}

			// Check all color fields are non-empty
			if theme.ChatBg == "" {
				t.Error("ChatBg is empty")
			}
			if theme.DiagramBg == "" {
				t.Error("DiagramBg is empty")
			}
			if theme.UserMsgBg == "" {
				t.Error("UserMsgBg is empty")
			}
			if theme.AgentMsgBg == "" {
				t.Error("AgentMsgBg is empty")
			}
			if theme.InputBg == "" {
				t.Error("InputBg is empty")
			}
			if theme.TextPrimary == "" {
				t.Error("TextPrimary is empty")
			}
			if theme.TextSecondary == "" {
				t.Error("TextSecondary is empty")
			}
			if theme.TextMuted == "" {
				t.Error("TextMuted is empty")
			}
			if theme.AccentUser == "" {
				t.Error("AccentUser is empty")
			}
			if theme.AccentAgent == "" {
				t.Error("AccentAgent is empty")
			}
			if theme.AccentCode == "" {
				t.Error("AccentCode is empty")
			}
		})
	}
}

func TestGetThemeNames(t *testing.T) {
	names := GetThemeNames()
	available := GetAvailableThemes()

	// Should return same list
	if len(names) != len(available) {
		t.Errorf("GetThemeNames() length = %d, GetAvailableThemes() length = %d", len(names), len(available))
	}

	for i := range names {
		if names[i] != available[i] {
			t.Errorf("GetThemeNames()[%d] = %q, GetAvailableThemes()[%d] = %q", i, names[i], i, available[i])
		}
	}
}

// Benchmark theme operations
func BenchmarkGetTheme(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetTheme("catppuccin-mocha")
	}
}

func BenchmarkSetActiveTheme(b *testing.B) {
	themes := []string{"catppuccin-mocha", "dracula", "nord"}

	for i := 0; i < b.N; i++ {
		SetActiveTheme(themes[i%len(themes)])
	}
}

func BenchmarkGetAvailableThemes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetAvailableThemes()
	}
}
