package writer

import (
	"bytes"
	"testing"
)

func TestNewBatchWriter(t *testing.T) {
	var buf bytes.Buffer

	// Create a new BatchWriter with a buffer size of 3
	bw := NewBatchWriter(3, &buf)
	// Write 2 strings to buffer to test if buffer is written
	err := bw.Write("test1")
	err = bw.Write("test2")
	if err != nil {
		t.Errorf("Error writing to file: %v", err)
	}
	if len(bw.buffer) != 2 {
		t.Errorf("Expected buffer length to be 2, got %d", len(bw.buffer))
	}

	// After 2 writes, should not have written to file yet
	actualFileContent := buf.String()
	if actualFileContent != "" {
		t.Errorf("Expected file contents to be '', got %s", actualFileContent)
	}

	// Write to buffer again to test if buffer size is 0 after flush
	err = bw.Write("test3")
	if err != nil {
		t.Errorf("Error writing to file: %v", err)
	}
	if len(bw.buffer) != 0 {
		t.Errorf("Expected buffer length to be 0, got %d", len(bw.buffer))
	}

	// Check if the file was written to output file correctly after flush
	actualFileContent = buf.String()
	if actualFileContent != "test1\ntest2\ntest3\n" {
		t.Errorf("Expected file contents to be 'test1\ntest2\ntest3\n', got %s", actualFileContent)
	}

}

func TestBatchWriter_Write(t *testing.T) {
	var buf bytes.Buffer
	// Create a new BatchWriter with a buffer size of 3
	bw := NewBatchWriter(3, &buf)

	// Write to buffer and check if buffer length is 1
	err := bw.Write("test1")
	if err != nil {
		t.Errorf("Error writing to file: %v", err)
	}
	if len(bw.buffer) != 1 {
		t.Errorf("Expected buffer length to be 1, got %d", len(bw.buffer))
	}

	// Write to buffer again to test if buffer size is 2
	err = bw.Write("test2")
	if err != nil {
		t.Errorf("Error writing to file: %v", err)
	}
	if len(bw.buffer) != 2 {
		t.Errorf("Expected buffer length to be 2, got %d", len(bw.buffer))
	}
}

func TestBatchWriter_Flush(t *testing.T) {
	var buf bytes.Buffer

	// Create a new BatchWriter with a buffer size of 3
	bw := NewBatchWriter(3, &buf)

	// Write to buffer
	err := bw.Write("test1")

	// Flush should write the buffer to the file
	err = bw.Flush()
	if err != nil {
		t.Errorf("Error flushing file: %v", err)
	}

	// Check if the file was written to output file correctly
	actualFileContent := buf.String()
	if actualFileContent != "test1\n" {
		t.Errorf("Expected file contents to be 'test1\n', got %s", actualFileContent)
	}
}

func TestBatchWriter_Close(t *testing.T) {
	var buf bytes.Buffer

	// Create a new BatchWriter with a buffer size of 3
	bw := NewBatchWriter(3, &buf)

	// Write to buffer
	err := bw.Write("test1")
	if err != nil {
		t.Errorf("Error writing to file: %v", err)
	}
	if len(bw.buffer) != 1 {
		t.Errorf("Expected buffer length to be 1, got %d", len(bw.buffer))
	}

	// Close should flush the buffer
	err = bw.Close()
	if err != nil {
		t.Errorf("Error closing file: %v", err)
	}

	// Check if the file was written to output file correctly
	actualFileContent := buf.String()
	if actualFileContent != "test1\n" {
		t.Errorf("Expected file contents to be 'test1\n', got %s", actualFileContent)
	}
}
