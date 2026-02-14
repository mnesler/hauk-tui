# Hauk TUI

A chat-based TUI for generating mermaid diagrams with LLM assistance.

## Features

- **Chat Interface**: Natural conversation with LLM agents
- **Dual Pane Layout**: Scrollable chat on the left, live application logs on the right
- **Mouse Support**: Scroll both panes independently with mouse wheel
- **Application Logs**: Real-time debug and event logging
- **Multiple LLM Providers**: OpenCode, OpenRouter, Anthropic, GitHub Copilot
- **Minimalist Design**: No borders, color-based panel separation
- **Theme Switching**: 8 color themes with live preview and persistent config

## Installation

```bash
go install github.com/mnesler/hauk-tui/cmd/hauk@latest
```

### Development Setup

After cloning the repository, install the pre-commit hook for build verification:

```bash
cp scripts/pre-commit .git/hooks/pre-commit
chmod +x .git/hooks/pre-commit
```

This hook automatically verifies that code compiles before allowing commits. See [scripts/README.md](scripts/README.md) for more details.

## Usage

```bash
hauk
```

### Keybindings

- `Enter` - Send message
- `Shift+Enter` - New line in message
- `Ctrl+C` or `Esc` - Quit

### Commands

- `/theme` - Open theme selector with live preview
  - Available themes: Catppuccin Mocha (default), Dracula, Nord, Gruvbox, Tokyo Night, GitHub Dark, Blue Monochrome Dark, Blue Monochrome
  - Use arrow keys (↑/↓) to preview themes in real-time
  - Press `Enter` to save selection, `Esc` to cancel

## Configuration

Configuration file location: `~/.config/hauk/config.yaml`

```yaml
llm:
  default_provider: opencode
  
ui:
  theme: catppuccin-mocha
```

## License

MIT License - See [LICENSE](LICENSE) for details.
