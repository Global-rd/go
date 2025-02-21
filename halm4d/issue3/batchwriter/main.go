package main

import (
	"batchwriter/client"
	"batchwriter/writer"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Setup batch writer
	outputFile := os.Getenv("OUTPUT_FILE")
	if outputFile == "" {
		fmt.Println("OUTPUT_FILE environment variable not set, applying default value: output.txt")
		outputFile = "output.txt"
	}

	reCreateOutputFile, err := strconv.ParseBool(os.Getenv("RECREATE_OUTPUT_FILE"))
	if err != nil {
		fmt.Println("Error parsing RECREATE_OUTPUT_FILE environment variable, applying default value:", err)
		reCreateOutputFile = true
	}

	maxBatchSize, err := strconv.Atoi(os.Getenv("MAX_BATCH_SIZE"))
	if err != nil {
		fmt.Println("Error parsing MAX_BATCH_SIZE environment variable, applying default value:", err)
		maxBatchSize = 3
	}

	batchWriter := writer.NewBatchWriter(maxBatchSize, outputFile, reCreateOutputFile)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			closeBatchWriter(batchWriter)
			os.Exit(1)
		}
		closeBatchWriter(batchWriter)
	}()

	// Setup panicking client
	serverUrl := os.Getenv("SERVER_URL")
	if serverUrl == "" {
		fmt.Println("SERVER_URL environment variable not set, applying default value: http://localhost:8080")
		serverUrl = "http://localhost:8080"
	}

	panicRate, err := strconv.Atoi(os.Getenv("PANIC_RATE"))
	if err != nil {
		fmt.Println("PANIC_RATE environment variable not set, applying default value: 10")
		panicRate = 10
	}
	panickingClient := client.NewPanickingClient(serverUrl)
	for i := 0; i < 10; i++ {
		bodyString, err := panickingClient.Get()
		if err != nil {
			log.Fatalln("Error getting response from server:", err)
		}

		requestCount, err := strconv.Atoi(strings.Split(bodyString, ": ")[1])
		if err != nil {
			log.Fatalln("Error parsing request count:", err)
		}

		if requestCount%panicRate == 0 {
			panic("Panic!")
		}
		err = batchWriter.Write(bodyString)
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}
	}

}

func closeBatchWriter(bw *writer.BatchWriter) {
	err := bw.Close()
	if err != nil {
		fmt.Println("Error closing batch writer:", err)
	}
	fmt.Println("Batch writer closed successfully. Lines written:", bw.WrittenLines())
}
