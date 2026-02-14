# Hauk TUI

A chat-based TUI for generating mermaid diagrams with LLM assistance.

## Features

- **Chat Interface**: Natural conversation with LLM agents
- **Dual Pane Layout**: Chat on the left, diagram preview on the right
- **Real-time Preview**: ASCII diagram rendering as you chat
- **Multiple LLM Providers**: OpenCode, OpenRouter, Anthropic, GitHub Copilot
- **Minimalist Design**: No borders, color-based panel separation
- **Theme Switching**: 6 color themes with live preview and persistent config

## Installation

```bash
go install github.com/mnesler/hauk-tui/cmd/hauk@latest
```

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
  - Available themes: Catppuccin Mocha (default), Dracula, Nord, Gruvbox, Tokyo Night, GitHub Dark
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
