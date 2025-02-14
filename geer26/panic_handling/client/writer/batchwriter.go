package writer

import (
	"errors"
	"log"
)

const (
	MAX_BUFFERSIZE = 15
)

var BufferFullError = errors.New("buffer full")

type Writer interface {
	Write(string) error
	Flush() error
	Close() error
}

type BatchWriter struct {
	buffer         []string
	max_buffersize int
}

// Panic occurs here, when buffer is full!
func (b *BatchWriter) Write(s string) error {
	if len(b.buffer) >= b.max_buffersize {
		return BufferFullError
	}
	b.buffer = append(b.buffer, s)
	log.Println(b.buffer)
	return nil
}

func (b *BatchWriter) Flush() error {
	return nil
}

func (b *BatchWriter) Close() error {
	return nil
}

func NewWriter() *BatchWriter {
	writer := BatchWriter{}
	writer.max_buffersize = MAX_BUFFERSIZE
	return &writer
}
