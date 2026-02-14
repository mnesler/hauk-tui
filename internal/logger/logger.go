package logger

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
)

// Logger is the main logger with circular buffer storage
type Logger struct {
	logger *logrus.Logger
	buffer *CircularBuffer
	mu     sync.Mutex
}

// Log is the global logger instance
var Log *Logger

// Level represents log levels
type Level = logrus.Level

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in production
	DebugLevel = logrus.DebugLevel
	// InfoLevel is the default logging priority
	InfoLevel = logrus.InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual human review
	WarnLevel = logrus.WarnLevel
	// ErrorLevel logs are high-priority
	ErrorLevel = logrus.ErrorLevel
)

// Fields type for structured logging
type Fields = logrus.Fields

// Init initializes the global logger with the specified buffer size
func Init(bufferSize int) {
	Log = NewLogger(bufferSize)
}

// NewLogger creates a new logger instance
func NewLogger(bufferSize int) *Logger {
	l := &Logger{
		logger: logrus.New(),
		buffer: NewCircularBuffer(bufferSize),
	}

	// Configure logrus
	l.logger.SetLevel(logrus.DebugLevel)
	l.logger.SetFormatter(&logrus.TextFormatter{
		DisableColors:   true,
		TimestampFormat: "15:04:05",
		FullTimestamp:   true,
	})

	// Set custom hook to capture logs in buffer
	l.logger.AddHook(l)

	return l
}

// Levels implements logrus.Hook interface
func (l *Logger) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire implements logrus.Hook interface
// This is called for every log entry and adds it to our circular buffer
func (l *Logger) Fire(entry *logrus.Entry) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Format the log entry
	formatted := formatLogEntry(entry)
	l.buffer.Add(formatted)

	return nil
}

// formatLogEntry formats a log entry for display
func formatLogEntry(entry *logrus.Entry) string {
	timestamp := entry.Time.Format("15:04:05")
	level := formatLevel(entry.Level)
	component := "app"

	// Extract component from fields if available
	if comp, ok := entry.Data["component"]; ok {
		component = fmt.Sprintf("%v", comp)
	}

	return fmt.Sprintf("[%s] %-5s [%s] %s", timestamp, level, component, entry.Message)
}

// formatLevel formats the log level as a string
func formatLevel(level logrus.Level) string {
	switch level {
	case logrus.DebugLevel:
		return "DEBUG"
	case logrus.InfoLevel:
		return "INFO"
	case logrus.WarnLevel:
		return "WARN"
	case logrus.ErrorLevel:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

// GetLogs returns all log entries in chronological order
func (l *Logger) GetLogs() []string {
	return l.buffer.GetAll()
}

// Debug logs a debug message
func (l *Logger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

// Debugf logs a formatted debug message
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

// Info logs an info message
func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

// Infof logs a formatted info message
func (l *Logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

// Warn logs a warning message
func (l *Logger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

// Warnf logs a formatted warning message
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

// Error logs an error message
func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

// Errorf logs a formatted error message
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

// WithField creates a new logger entry with a single field
func (l *Logger) WithField(key string, value interface{}) *logrus.Entry {
	return l.logger.WithField(key, value)
}

// WithFields creates a new logger entry with multiple fields
func (l *Logger) WithFields(fields Fields) *logrus.Entry {
	return l.logger.WithFields(fields)
}

// GetLogCount returns the number of log entries in the buffer
func (l *Logger) GetLogCount() int {
	return l.buffer.Count()
}

// Clear clears all log entries from the buffer
func (l *Logger) Clear() {
	l.buffer.Clear()
}

// GetLogs returns all log entries from the global logger
func GetLogs() []string {
	if Log == nil {
		return []string{"Logger not initialized"}
	}
	return Log.GetLogs()
}

// Component creates a logger with a component field
func Component(name string) *logrus.Entry {
	if Log == nil {
		// Return a dummy entry if logger not initialized
		l := logrus.New()
		l.Out = nil // Discard output
		return l.WithField("component", name)
	}
	return Log.WithField("component", name)
}

// StartupMessage logs the application startup banner
func StartupMessage(version string) {
	if Log == nil {
		return
	}

	Log.Info("════════════════════════════════════════")
	Log.Infof("    Hauk-TUI %s", version)
	Log.Info("    Terminal User Interface")
	Log.Info("════════════════════════════════════════")
	Component("app").Info("Application starting...")
}
