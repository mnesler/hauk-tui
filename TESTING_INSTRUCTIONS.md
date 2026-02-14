# Testing Instructions for Blue Monochrome Themes

## Manual Testing Required

Since Go is not available in the current environment, the following tests need to be run manually:

### 1. Run Automated Tests
```bash
# Run all tests with coverage
go test ./... -cover

# Run tests with race detector
go test ./... -race

# Generate coverage report
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

**Expected Results**:
- All tests should pass
- Coverage should remain above 50%
- No race conditions detected
- Test count should increase from previous runs (2 new themes = ~10+ new test cases)

### 2. Build the Binary
```bash
go build -o bin/hauk cmd/hauk/main.go
```

**Expected**: Binary should build without errors

### 3. Manual Theme Testing

Run the application and test both new themes:

```bash
./bin/hauk
```

#### Test Blue Monochrome Dark Theme
1. Type `/theme` and press Enter
2. Use arrow keys to navigate to "Blue Monochrome Dark"
3. **Verify live preview**:
   - Background should be very dark blue (#0041C7)
   - Text should be bright cyan (#3ACBE8)
   - User messages should have medium blue background (#0D85D8)
   - Agent messages should have darker blue background (#0160C9)
   - All colors should be from the blue palette (no greens, purples, etc.)
4. Press Enter to select
5. Quit and restart the app
6. **Verify persistence**: Theme should still be Blue Monochrome Dark

#### Test Blue Monochrome Theme
1. Type `/theme` and press Enter
2. Use arrow keys to navigate to "Blue Monochrome"
3. **Verify live preview**:
   - Background should be darker blue (#0160C9)
   - Text should be white (#FFFFFF) for maximum readability
   - User messages should have bright blue background (#1CA3DE)
   - Agent messages should have medium blue background (#0D85D8)
   - Accents should be bright blue
4. Press Enter to select
5. Quit and restart the app
6. **Verify persistence**: Theme should still be Blue Monochrome

#### Test Theme Switching
1. Switch between the two blue themes multiple times
2. Switch from blue themes to other themes (e.g., Dracula, Nord)
3. Switch back to blue themes
4. Press Esc to cancel theme selection (should revert to previous theme)

### 4. Config File Verification

Check that the config file saves correctly:

```bash
cat ~/.config/hauk/config.yaml
```

**Expected content examples**:
```yaml
llm:
  default_provider: opencode
ui:
  theme: blue-monochrome-dark
```

or

```yaml
llm:
  default_provider: opencode
ui:
  theme: blue-monochrome
```

### 5. Visual Comparison

Compare the themes against the preview document:

```bash
cat BLUE_THEME_PREVIEW.md
```

Verify that the actual rendered colors match the specifications in the preview.

## Files Modified

The following files were modified to add the blue monochrome themes:

1. **internal/ui/theme.go** - Added `BlueMonochromeDark` and `BlueMonochrome` theme definitions
2. **internal/ui/theme_registry.go** - Registered both themes in `AvailableThemes` map
3. **internal/app/update.go** - Added display names for both themes in `formatThemeName()`
4. **internal/ui/theme_test.go** - Updated test cases to include both new themes (6→8 themes)
5. **internal/app/update_test.go** - Added formatThemeName test cases for both new themes
6. **README.md** - Updated documentation to list new themes

## Expected Test Coverage

- **Overall coverage**: Should remain 50%+ (currently 51.1%)
- **Theme registry tests**: Should now test 8 themes instead of 6
- **Command parser tests**: No changes needed (still 90.9%)
- **Config tests**: No changes needed (still 75.9%)
- **App update tests**: Should include new themes in formatThemeName tests

## CI/CD

Once changes are pushed, GitHub Actions will automatically:
1. Run all tests with race detector
2. Build the binary for Linux, macOS, and Windows
3. Check code coverage threshold (40% minimum)
4. Run linters (golangci-lint)

Monitor the workflow at: https://github.com/mnesler/hauk-tui/actions

## Troubleshooting

### If tests fail:
- Check that all theme names are lowercase with hyphens (not camelCase)
- Verify hex colors are uppercase (e.g., `#3ACBE8` not `#3acbe8`)
- Ensure all 11 theme properties are set for both themes

### If themes don't display correctly:
- Check terminal color support: `echo $TERM`
- Try different terminal emulators (iTerm2, Alacritty, Windows Terminal)
- Verify true color support: Most modern terminals support 24-bit color

### If config doesn't persist:
- Check file permissions: `ls -la ~/.config/hauk/`
- Manually create directory if needed: `mkdir -p ~/.config/hauk`
- Check for write errors in application logs
