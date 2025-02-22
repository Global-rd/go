package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
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
        return bw.Flush()
    }
    return nil
}

func (bw *BatchWriter) Flush() error {
    _, err := bw.target.Write([]byte(bw.buffer.String()))
    bw.buffer.Reset()
    return err
}

func (bw *BatchWriter) Close() error {
    err := bw.Flush()
    if closer, ok := bw.target.(io.Closer); ok {
        return closer.Close()
    }
    return err
}

func processData(source io.Reader, writer Writer) error {
    scanner := bufio.NewScanner(source)
    for scanner.Scan() {
        data := scanner.Text()
        if err := writer.Write(data + "\n"); err != nil {
            return err
        }
    }
    return scanner.Err()
}

func main() {
    sourceFile, err := os.Open("source.txt")
    if err != nil {
        fmt.Println("Error opening source file:", err)
        return
    }
    defer sourceFile.Close()

    targetFile, err := os.Create("target.txt")
    if err != nil {
        fmt.Println("Error creating target file:", err)
        return
    }
    defer targetFile.Close()

    batchWriter := NewBatchWriter(targetFile, 1024) // 1KB batch size
    defer batchWriter.Close()

    err = processDataWithRecovery(sourceFile, batchWriter)
    if err != nil {
        fmt.Println("Error processing data:", err)
    }
}

func processDataWithRecovery(source io.Reader, writer Writer) (err error) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Panic occurred:", r)
            err = fmt.Errorf("panic: %v", r)
            // Ensure buffered data is flushed on panic
            flushErr := writer.Flush()
            if flushErr != nil {
                fmt.Println("Error flushing buffer:", flushErr)
            }
        }
    }()

    return processData(source, writer)
}
