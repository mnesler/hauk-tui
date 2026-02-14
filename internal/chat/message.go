package chat

import "time"

// Role represents who sent a message
type Role string

const (
	RoleUser  Role = "user"
	RoleAgent Role = "agent"
)

// Message represents a single chat message
type Message struct {
	Role      Role
	Content   string
	Timestamp time.Time
	Diagram   string // Optional extracted mermaid code
}

// NewMessage creates a new message
func NewMessage(role Role, content string) Message {
	return Message{
		Role:      role,
		Content:   content,
		Timestamp: time.Now(),
	}
}

// HasDiagram returns true if this message contains a diagram
func (m Message) HasDiagram() bool {
	return m.Diagram != ""
}
