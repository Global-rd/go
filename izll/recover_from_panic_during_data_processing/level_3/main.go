package level_3

import (
	"io"
	"net/http"
	"os"
	"recover_from_panic_during_data_processing/writer"
)

const sourceURL = "https://api.artic.edu/api/v1/artworks?limit=20"
const targetFilePath = "/tmp/output.txt"

func Run() {
	targetFile, err := os.Create(targetFilePath)
	bw := writer.NewBatchWriter(targetFile, 100)
	defer bw.Close()

	client := &http.Client{}
	resp, err := client.Get(sourceURL)
	if err != nil {
		println("Error getting data from source")
		return
	}
	defer resp.Body.Close()

	buffer := make([]byte, 1024)
	hasError := false

	println("Reading source url:", sourceURL)

	for {
		n, err := resp.Body.Read(buffer)

		if n > 0 {
			writeErr := bw.Write(string(buffer[:n]))
			if writeErr != nil {
				hasError = true
				println("Error writing data:", writeErr.Error())
				break
			}
		}

		if err != nil {
			if err != io.EOF {
				println("Error reading data:", err.Error())
				hasError = true
			}
			break
		}
	}

	if !hasError {
		println("Data processing completed")
	}
}
