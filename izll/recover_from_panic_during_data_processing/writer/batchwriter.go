package writer

import (
	"io"
	"strings"
)

type Writer interface {
	Write(string) error
	Flush() error
	Close() error
}

type BatchWriter struct {
	buffer    *strings.Builder
	target    io.Writer
	batchSize int
}

func NewBatchWriter(target io.Writer, batchSize int) *BatchWriter {
	return &BatchWriter{
		buffer:    &strings.Builder{},
		target:    target,
		batchSize: batchSize,
	}
}

func (bw *BatchWriter) Write(data string) error {
	bw.buffer.WriteString(data)
	if bw.buffer.Len() >= bw.batchSize {
		err := bw.Flush()
		if err != nil {
			return err
		}
	}
	return nil
}

func (bw *BatchWriter) Flush() error {
	_, err := bw.target.Write([]byte(bw.buffer.String()))
	if err != nil {
		return err
	}
	bw.buffer.Reset()
	return nil
}

func (bw *BatchWriter) Close() error {
	err := bw.Flush()
	if err != nil {
		return err
	}
	return nil
}
