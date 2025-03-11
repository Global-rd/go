package level_2

import (
	"io"
	"os"
	"recover_from_panic_during_data_processing/writer"
)

const sourceFilePath = "level_2/test_file.json"
const targetFilePath = "/tmp/output.txt"

func Run() {
	targetFile, err := os.Create(targetFilePath)
	if err != nil {
		println("Error creating target file")
		return
	}
	defer targetFile.Close()

	sourceFile, err := os.Open(sourceFilePath)
	if err != nil {
		println("Error opening source file")
		return
	}
	defer sourceFile.Close()

	bw := writer.NewBatchWriter(targetFile, 100)
	defer bw.Close()

	hasError := false
	buffer := make([]byte, 100)

	println("Reading source file")

	for {
		n, err := sourceFile.Read(buffer)
		if n > 0 {
			// Ha van olvasott adat, írjuk ki
			writeErr := bw.Write(string(buffer[:n]))
			if writeErr != nil {
				hasError = true
				println("Error writing data")
				break
			}
		}

		if err != nil {
			if err != io.EOF {
				hasError = true
				println("Error reading data")
			}
			break
		}
	}

	if !hasError {
		println("Data processing completed")
	}
}
