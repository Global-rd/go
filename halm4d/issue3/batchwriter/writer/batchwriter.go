package writer

import (
	"log"
	"os"
)

type BatchWriter struct {
	buffer       []string
	maxBatchSize int
	outputFile   string
	writtenLines int
}

func removeAndCreateFile(outputFile string) error {
	if _, err := os.Stat(outputFile); err == nil {
		err := os.Remove(outputFile)
		if err != nil {
			return err
		}
	}
	file, err := os.Create(outputFile)
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	if err != nil {
		return err
	}
	return err
}

func NewBatchWriter(maxBatchSize int, outputFile string, reCreateOutputFile bool) *BatchWriter {
	if reCreateOutputFile {
		err := removeAndCreateFile(outputFile)
		if err != nil {
			log.Fatalln("Error recreating output file:", err)
		}
	}
	return &BatchWriter{
		buffer:       make([]string, 0),
		maxBatchSize: maxBatchSize,
		outputFile:   outputFile,
	}
}

func (bw *BatchWriter) Write(s string) error {
	bw.buffer = append(bw.buffer, s)
	if len(bw.buffer) >= bw.maxBatchSize {
		return bw.Flush()
	}
	return nil
}

func (bw *BatchWriter) Flush() error {
	file, err := os.OpenFile(bw.outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 06666)
	if err != nil {
		return err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	for _, s := range bw.buffer {
		_, err = file.WriteString(s + "\n")
		if err != nil {
			log.Fatalln(err)
		}
		bw.writtenLines++
	}
	bw.buffer = make([]string, 0)
	return nil
}

func (bw *BatchWriter) Close() error {
	return bw.Flush()
}

func (bw *BatchWriter) WrittenLines() int {
	return bw.writtenLines
}
