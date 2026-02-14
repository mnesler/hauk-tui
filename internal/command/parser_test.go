package command

import (
	"testing"
)

func TestParseCommand(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantCmd  CommandType
		wantArgs []string
	}{
		{
			name:     "valid theme command",
			input:    "/theme",
			wantCmd:  CommandTheme,
			wantArgs: nil,
		},
		{
			name:     "theme command with leading whitespace",
			input:    "  /theme",
			wantCmd:  CommandTheme,
			wantArgs: nil,
		},
		{
			name:     "theme command with trailing whitespace",
			input:    "/theme  ",
			wantCmd:  CommandTheme,
			wantArgs: nil,
		},
		{
			name:     "theme command uppercase",
			input:    "/THEME",
			wantCmd:  CommandTheme,
			wantArgs: nil,
		},
		{
			name:     "theme command mixed case",
			input:    "/ThEmE",
			wantCmd:  CommandTheme,
			wantArgs: nil,
		},
		{
			name:     "theme command with args",
			input:    "/theme dark",
			wantCmd:  CommandTheme,
			wantArgs: []string{"dark"},
		},
		{
			name:     "theme command with multiple args",
			input:    "/theme dark nord",
			wantCmd:  CommandTheme,
			wantArgs: []string{"dark", "nord"},
		},
		{
			name:     "invalid command",
			input:    "/invalid",
			wantCmd:  CommandNone,
			wantArgs: nil,
		},
		{
			name:     "no slash prefix",
			input:    "theme",
			wantCmd:  CommandNone,
			wantArgs: nil,
		},
		{
			name:     "empty string",
			input:    "",
			wantCmd:  CommandNone,
			wantArgs: nil,
		},
		{
			name:     "just slash",
			input:    "/",
			wantCmd:  CommandNone,
			wantArgs: nil,
		},
		{
			name:     "only whitespace",
			input:    "   ",
			wantCmd:  CommandNone,
			wantArgs: nil,
		},
		{
			name:     "regular message",
			input:    "hello world",
			wantCmd:  CommandNone,
			wantArgs: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCmd, gotArgs := ParseCommand(tt.input)

			if gotCmd != tt.wantCmd {
				t.Errorf("ParseCommand() gotCmd = %v, want %v", gotCmd, tt.wantCmd)
			}

			// Compare args
			if len(gotArgs) != len(tt.wantArgs) {
				t.Errorf("ParseCommand() gotArgs length = %v, want %v", len(gotArgs), len(tt.wantArgs))
				return
			}

			for i := range gotArgs {
				if gotArgs[i] != tt.wantArgs[i] {
					t.Errorf("ParseCommand() gotArgs[%d] = %v, want %v", i, gotArgs[i], tt.wantArgs[i])
				}
			}
		})
	}
}

func TestCommandType_String(t *testing.T) {
	tests := []struct {
		name string
		cmd  CommandType
		want string
	}{
		{
			name: "CommandNone",
			cmd:  CommandNone,
			want: "none",
		},
		{
			name: "CommandTheme",
			cmd:  CommandTheme,
			want: "theme",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// For now, just check the value is correct
			// We can add String() method later if needed
			if tt.cmd == CommandNone && int(tt.cmd) != 0 {
				t.Errorf("CommandNone should be 0, got %d", tt.cmd)
			}
			if tt.cmd == CommandTheme && int(tt.cmd) != 1 {
				t.Errorf("CommandTheme should be 1, got %d", tt.cmd)
			}
		})
	}
}

// Benchmark command parsing
func BenchmarkParseCommand(b *testing.B) {
	inputs := []string{
		"/theme",
		"/theme dark",
		"regular message",
		"/invalid",
	}

	for _, input := range inputs {
		b.Run(input, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ParseCommand(input)
			}
		})
	}
}
