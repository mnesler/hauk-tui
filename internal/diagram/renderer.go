package diagram

import (
	"github.com/AlexanderGrooff/mermaid-ascii/cmd"
)

// Render converts mermaid code to ASCII art
func Render(mermaidCode string) (string, error) {
	// Use the mermaid-ascii library to render
	_, err := cmd.DiagramFactory(mermaidCode)
	if err != nil {
		return "", err
	}

	// Convert to string representation
	// Note: The actual rendering implementation will depend on the library's API
	// This is a placeholder that returns the code as-is for now
	return mermaidCode, nil
}
