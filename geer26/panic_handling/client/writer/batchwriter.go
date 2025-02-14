package writer

import (
	"errors"
	"fmt"
	"log"
	"os"
)

const (
	MAX_BUFFERSIZE = 15
	FILEPATH       = "./results.txt"
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
	return nil
}

func (b *BatchWriter) Flush() error {

	f, err := os.Create(FILEPATH)
	if err != nil {
		return err
	}
	defer f.Close()

	for i, entry := range b.buffer {
		_, err := f.Write([]byte(fmt.Sprintf("#%d ENTRY: %s\n", i+1, entry)))
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *BatchWriter) Close() error {
	err := b.Flush()
	return err
}

func (b *BatchWriter) Display() {
	log.Println(b.buffer)
}

func NewWriter() *BatchWriter {
	writer := BatchWriter{}
	writer.max_buffersize = MAX_BUFFERSIZE
	return &writer
}
