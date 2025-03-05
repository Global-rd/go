package writer

import (
	"errors"
	"fmt"
	"os"
)

const (
	MAX_BUFFERSIZE = 15
	FILEPATH       = "./results.txt"
	// Simon says, any string that contains this character is invalid, and
	// should stop the app, but the collected data must be saved into the
	// results.txt file
	INVALID_CHAR = "A"
)

// Custom error declarations
var BufferFullError = errors.New("buffer full")
var ContainsInvalidCharacterError = errors.New("invalid character")

// Predefined interface declarations
type Writer interface {
	Write(string) error
	Flush() error
	Close() error
}

type BatchWriter struct {
	// simple buffer, slice of strings
	buffer         []string
	max_buffersize int
}

// Error occurs here, when buffer is full or in case of invalid string!
func (b *BatchWriter) Write(s string) error {
	if !CheckIfValid(s) {
		return ContainsInvalidCharacterError
	}

	if len(b.buffer) >= b.max_buffersize {
		return BufferFullError
	}
	// If everything seems ok, do not return error!
	b.buffer = append(b.buffer, s)
	return nil
}

func (b *BatchWriter) Flush() error {

	// Opens a file (creates, if doe not exist)
	f, err := os.Create(FILEPATH)
	if err != nil {
		return err
	}
	defer f.Close()

	// flush the buffer into the file
	for i, entry := range b.buffer {
		_, err := f.WriteString(fmt.Sprintf("#%d ENTRY: %s\n", i+1, entry))
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *BatchWriter) Close() error {
	// Flushes the buffer into the file anyway
	err := b.Flush()
	return err
}

// Constructor for the BatchWriter
func NewWriter() *BatchWriter {
	writer := BatchWriter{}
	writer.max_buffersize = MAX_BUFFERSIZE
	return &writer
}

func CheckIfValid(s string) bool {
	// If passed String has an appercase A, return false
	// as is invalid!
	for _, b := range s {
		if string(b) == INVALID_CHAR {
			return false
		}
	}
	return true
}
