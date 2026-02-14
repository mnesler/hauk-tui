package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	
	if cfg == nil {
		t.Fatal("DefaultConfig() returned nil")
	}
	
	if cfg.Theme != "catppuccin-mocha" {
		t.Errorf("DefaultConfig().Theme = %q, want %q", cfg.Theme, "catppuccin-mocha")
	}
}

func TestConfigPath(t *testing.T) {
	path, err := ConfigPath()
	if err != nil {
		t.Fatalf("ConfigPath() error = %v", err)
	}
	
	if path == "" {
		t.Error("ConfigPath() returned empty string")
	}
	
	// Should contain .config/hauk/config.yaml
	if !filepath.IsAbs(path) {
		t.Error("ConfigPath() should return absolute path")
	}
	
	// Check that path ends with expected suffix
	if !strings.HasSuffix(path, filepath.Join(".config", "hauk", "config.yaml")) {
		t.Errorf("ConfigPath() = %q, should end with .config/hauk/config.yaml", path)
	}
}

func TestLoad_NoFile(t *testing.T) {
	// Test loading when file doesn't exist
	// Should return default config without error
	
	// Save original home
	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)
	
	// Use temp directory
	tempDir := t.TempDir()
	os.Setenv("HOME", tempDir)
	
	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() with no file error = %v, want nil", err)
	}
	
	if cfg == nil {
		t.Fatal("Load() returned nil config")
	}
	
	// Should return default
	if cfg.Theme != "catppuccin-mocha" {
		t.Errorf("Load() with no file theme = %q, want %q", cfg.Theme, "catppuccin-mocha")
	}
}

func TestSaveAndLoad(t *testing.T) {
	// Test saving and loading a config
	
	// Save original home
	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)
	
	// Use temp directory
	tempDir := t.TempDir()
	os.Setenv("HOME", tempDir)
	
	// Create a config
	cfg := &Config{
		Theme: "dracula",
	}
	
	// Save it
	err := Save(cfg)
	if err != nil {
		t.Fatalf("Save() error = %v", err)
	}
	
	// Verify file was created
	path, _ := ConfigPath()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Fatal("Config file was not created")
	}
	
	// Load it back
	loaded, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	
	if loaded.Theme != "dracula" {
		t.Errorf("Loaded theme = %q, want %q", loaded.Theme, "dracula")
	}
}

func TestSave_CreatesDirectory(t *testing.T) {
	// Test that Save creates the directory if it doesn't exist
	
	// Save original home
	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)
	
	// Use temp directory
	tempDir := t.TempDir()
	os.Setenv("HOME", tempDir)
	
	cfg := DefaultConfig()
	
	// Save should create directory
	err := Save(cfg)
	if err != nil {
		t.Fatalf("Save() error = %v", err)
	}
	
	// Check directory exists
	configDir := filepath.Join(tempDir, ".config", "hauk")
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		t.Error("Save() did not create config directory")
	}
}

func TestLoad_InvalidYAML(t *testing.T) {
	// Test loading invalid YAML
	
	// Save original home
	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)
	
	// Use temp directory
	tempDir := t.TempDir()
	os.Setenv("HOME", tempDir)
	
	// Create config directory
	configDir := filepath.Join(tempDir, ".config", "hauk")
	err := os.MkdirAll(configDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create config dir: %v", err)
	}
	
	// Write invalid YAML
	configPath := filepath.Join(configDir, "config.yaml")
	invalidYAML := "theme: [invalid yaml"
	err = os.WriteFile(configPath, []byte(invalidYAML), 0644)
	if err != nil {
		t.Fatalf("Failed to write invalid YAML: %v", err)
	}
	
	// Load should return error
	_, err = Load()
	if err == nil {
		t.Error("Load() with invalid YAML should return error")
	}
}

func TestSave_RoundTrip(t *testing.T) {
	// Test multiple save/load cycles
	
	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)
	
	tempDir := t.TempDir()
	os.Setenv("HOME", tempDir)
	
	themes := []string{"catppuccin-mocha", "dracula", "nord", "gruvbox", "tokyo-night", "github-dark"}
	
	for _, theme := range themes {
		t.Run(theme, func(t *testing.T) {
			cfg := &Config{Theme: theme}
			
			if err := Save(cfg); err != nil {
				t.Fatalf("Save(%q) error = %v", theme, err)
			}
			
			loaded, err := Load()
			if err != nil {
				t.Fatalf("Load() after Save(%q) error = %v", theme, err)
			}
			
			if loaded.Theme != theme {
				t.Errorf("Round trip failed: saved %q, loaded %q", theme, loaded.Theme)
			}
		})
	}
}

// Benchmark config operations
func BenchmarkLoad(b *testing.B) {
	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)
	
	tempDir := b.TempDir()
	os.Setenv("HOME", tempDir)
	
	// Create a config file
	cfg := DefaultConfig()
	_ = Save(cfg)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Load()
	}
}

func BenchmarkSave(b *testing.B) {
	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)
	
	tempDir := b.TempDir()
	os.Setenv("HOME", tempDir)
	
	cfg := DefaultConfig()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Save(cfg)
	}
}
