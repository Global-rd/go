package writer

import (
	"log"
	"os"
)

type BatchWriter struct {
	buffer       []string
	maxBatchSize int
	outputFile   *os.File
	writtenLines int
}

func removeAndCreateFile(outputFile string) (file *os.File, err error) {
	if _, err := os.Stat(outputFile); err == nil {
		err := os.Remove(outputFile)
		if err != nil {
			return nil, err
		}
	}
	file, err = os.Create(outputFile)
	if err != nil {
		return file, err
	}
	return file, err
}

func NewBatchWriter(maxBatchSize int, outputFilePath string, reCreateOutputFile bool) *BatchWriter {
	var outputFile *os.File
	if reCreateOutputFile {
		var err error
		outputFile, err = removeAndCreateFile(outputFilePath)
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

func (bw *BatchWriter) Flush() (err error) {
	file, err := os.OpenFile(bw.outputFile.Name(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 06666)
	if err != nil {
		return err
	}
	defer func() {
		err = file.Close()
	}()
	for _, s := range bw.buffer {
		_, err = file.WriteString(s + "\n")
		if err != nil {
			return err
		}
		bw.writtenLines++
	}
	bw.buffer = make([]string, 0)
	return err
}

func (bw *BatchWriter) Close() error {
	return bw.Flush()
}

func (bw *BatchWriter) WrittenLines() int {
	return bw.writtenLines
}
