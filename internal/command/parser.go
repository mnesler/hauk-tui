package command

import "strings"

// CommandType represents a slash command
type CommandType int

const (
	CommandNone CommandType = iota
	CommandTheme
	// Future commands can be added here
)

// ParseCommand detects and parses slash commands from user input
func ParseCommand(input string) (CommandType, []string) {
	input = strings.TrimSpace(input)
	
	// Check if it starts with a slash
	if !strings.HasPrefix(input, "/") {
		return CommandNone, nil
	}
	
	// Split command and arguments
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return CommandNone, nil
	}
	
	// Get command (without the slash)
	cmd := strings.ToLower(strings.TrimPrefix(parts[0], "/"))
	args := parts[1:]
	
	// Match command
	switch cmd {
	case "theme":
		return CommandTheme, args
	default:
		return CommandNone, nil
	}
}
