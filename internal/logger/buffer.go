package logger

import (
	"sync"
)

// CircularBuffer is a thread-safe circular buffer for storing log entries
type CircularBuffer struct {
	entries []string
	size    int
	head    int
	count   int
	mu      sync.RWMutex
}

// NewCircularBuffer creates a new circular buffer with the given size
func NewCircularBuffer(size int) *CircularBuffer {
	return &CircularBuffer{
		entries: make([]string, size),
		size:    size,
		head:    0,
		count:   0,
	}
}

// Add adds a new entry to the buffer
// If the buffer is full, it overwrites the oldest entry (FIFO)
func (b *CircularBuffer) Add(entry string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.entries[b.head] = entry
	b.head = (b.head + 1) % b.size

	if b.count < b.size {
		b.count++
	}
}

// GetAll returns all entries in chronological order (oldest to newest)
func (b *CircularBuffer) GetAll() []string {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if b.count == 0 {
		return []string{}
	}

	result := make([]string, b.count)

	// If buffer is not full, entries are from 0 to count-1
	if b.count < b.size {
		copy(result, b.entries[:b.count])
		return result
	}

	// If buffer is full, entries are split around head
	// Copy from head to end (oldest entries)
	oldestCount := b.size - b.head
	copy(result, b.entries[b.head:])

	// Copy from start to head (newest entries)
	copy(result[oldestCount:], b.entries[:b.head])

	return result
}

// Count returns the number of entries in the buffer
func (b *CircularBuffer) Count() int {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.count
}

// Clear removes all entries from the buffer
func (b *CircularBuffer) Clear() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.head = 0
	b.count = 0
	b.entries = make([]string, b.size)
}
