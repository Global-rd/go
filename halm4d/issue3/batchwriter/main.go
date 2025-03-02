package main

import (
	"batchwriter/client"
	"batchwriter/config"
	"batchwriter/writer"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Setup batch writer
	conf := config.NewConfig()

	var err error
	var outputFile *os.File
	if conf.ReCreateOutputFile() {
		outputFile, err = removeAndCreateFile(conf.OutputFileName())
		if err != nil {
			log.Fatalln("Error re creating file:", err)
		}
	} else {
		outputFile, err = os.OpenFile(conf.OutputFileName(), os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalln("Error opening file:", err)
		}
	}

	defer func() {
		err = outputFile.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		} else {
			fmt.Println("Output file closed successfully.")
		}
	}()

	batchWriter := writer.NewBatchWriter(conf.MaxBatchSize(), outputFile)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
		if batchWriter != nil {
			err = batchWriter.Close()
			if err != nil {
				fmt.Println("Error closing batch writer:", err)
			} else {
				fmt.Println("Batch writer closed successfully. Lines written:", batchWriter.WrittenLines())
			}
		}
	}()

	// Setup panicking client
	panickingClient := client.NewPanickingClient(conf.ServerUrl())
	for i := 0; i < 10; i++ {
		bodyString, err := panickingClient.Get()
		if err != nil {
			log.Fatalln("Error getting response from server:", err)
		}

		requestCount, err := strconv.Atoi(strings.Split(bodyString, ": ")[1])
		if err != nil {
			log.Fatalln("Error parsing request count:", err)
		}

		if requestCount%conf.PanicRate() == 0 {
			panic("Panic!")
		}
		err = batchWriter.Write(bodyString)
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}
	}

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
