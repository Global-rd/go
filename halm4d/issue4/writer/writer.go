package writer

import (
	"encoding/csv"
	"io"
)

type Writer interface {
	Write([]byte) error
	Flush() error
	Close() error
}

type CSVWriter struct {
	writer *csv.Writer
}

func NewCSVConverter(output io.Writer) *CSVWriter {
	csvWriter := csv.NewWriter(output)
	csvWriter.Comma = ';'
	return &CSVWriter{
		writer: csvWriter,
	}
}

func (c *CSVWriter) Write(data [][]string) error {
	for _, v := range data {
		err := c.writer.Write(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *CSVWriter) Flush() error {
	c.writer.Flush()
	return c.writer.Error()
}

func (c *CSVWriter) Close() error {
	return c.Flush()
}
