package writer

import (
	"io"
)

type BatchWriter struct {
	buffer       []string
	maxBatchSize int
	writer       io.Writer
	writtenLines int
}

func NewBatchWriter(maxBatchSize int, writer io.Writer) *BatchWriter {
	return &BatchWriter{
		buffer:       make([]string, 0),
		maxBatchSize: maxBatchSize,
		writer:       writer,
	}
}

func (bw *BatchWriter) Write(s string) error {
	bw.buffer = append(bw.buffer, s)
	if len(bw.buffer) >= bw.maxBatchSize {
		return bw.Flush()
	}
	return nil
}

func (bw *BatchWriter) Flush() (err error) {
	for _, s := range bw.buffer {
		_, err = bw.writer.Write([]byte(s + "\n"))
		if err != nil {
			return err
		}
		bw.writtenLines++
	}
	bw.buffer = make([]string, 0)
	return nil
}

func (bw *BatchWriter) Close() (err error) {
	return bw.Flush()
}

func (bw *BatchWriter) WrittenLines() int {
	return bw.writtenLines
}
