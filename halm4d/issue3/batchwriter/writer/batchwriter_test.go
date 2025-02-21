package writer

import (
	"io"
	"os"
	"testing"
)

func TestNewBatchWriter(t *testing.T) {
	// Create a temp file to write to and remove it after the test is done
	temp := createTempFile(t)
	defer func() {
		err := os.Remove(temp.Name())
		if err != nil {
			t.Errorf("Error removing temp file: %v", err)
		}
	}()

	// Create a new BatchWriter with a buffer size of 3
	bw := NewBatchWriter(3, temp.Name(), true)
	defer func() {
		err := bw.outputFile.Close()
		if err != nil {
			t.Errorf("Error closing temp file: %v", err)
		}
	}()

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
	actualFileContent := readFile(t, temp.Name())
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
	actualFileContent = readFile(t, temp.Name())
	if actualFileContent != "test1\ntest2\ntest3\n" {
		t.Errorf("Expected file contents to be 'test1\ntest2\ntest3\n', got %s", actualFileContent)
	}

}

func TestBatchWriter_Write(t *testing.T) {
	// Create a temp file to write to and remove it after the test is done
	temp := createTempFile(t)
	defer func() {
		err := os.Remove(temp.Name())
		if err != nil {
			t.Errorf("Error removing temp file: %v", err)
		}
	}()

	// Create a new BatchWriter with a buffer size of 3
	bw := NewBatchWriter(3, temp.Name(), true)
	defer func() {
		err := bw.outputFile.Close()
		if err != nil {
			t.Errorf("Error closing temp file: %v", err)
		}
	}()

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
	// Create a temp file to write to and remove it after the test is done
	temp := createTempFile(t)
	defer func() {
		err := os.Remove(temp.Name())
		if err != nil {
			t.Errorf("Error removing temp file: %v", err)
		}
	}()

	// Create a new BatchWriter with a buffer size of 3
	bw := NewBatchWriter(3, temp.Name(), true)
	defer func() {
		err := bw.outputFile.Close()
		if err != nil {
			t.Errorf("Error closing temp file: %v", err)
		}
	}()

	// Write to buffer
	err := bw.Write("test1")

	// Flush should write the buffer to the file
	err = bw.Flush()
	if err != nil {
		t.Errorf("Error flushing file: %v", err)
	}

	// Check if the file was written to output file correctly
	actualFileContent := readFile(t, temp.Name())
	if actualFileContent != "test1\n" {
		t.Errorf("Expected file contents to be 'test1\n', got %s", actualFileContent)
	}
}

func TestBatchWriter_Close(t *testing.T) {
	// Create a temp file to write to and remove it after the test is done
	temp := createTempFile(t)
	defer func() {
		err := os.Remove(temp.Name())
		if err != nil {
			t.Errorf("Error removing temp file: %v", err)
		}
	}()

	// Create a new BatchWriter with a buffer size of 3
	bw := NewBatchWriter(3, temp.Name(), true)
	defer func() {
		err := bw.outputFile.Close()
		if err != nil {
			t.Errorf("Error closing temp file: %v", err)
		}
	}()

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
	actualFileContent := readFile(t, temp.Name())
	if actualFileContent != "test1\n" {
		t.Errorf("Expected file contents to be 'test1\n', got %s", actualFileContent)
	}
}

func readFile(t *testing.T, filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		t.Errorf("Error opening file: %v", err)
	}
	all, err := io.ReadAll(file)
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			t.Errorf("Error closing file: %v", err)
		}
	}()
	return string(all)
}

func createTempFile(t *testing.T) *os.File {
	temp, err := os.CreateTemp("", "output.txt")
	if err != nil {
		t.Errorf("Error creating temp file: %v", err)
	}
	defer func() {
		err := temp.Close()
		if err != nil {
			t.Errorf("Error closing temp file: %v", err)
		}
	}()
	return temp
}
