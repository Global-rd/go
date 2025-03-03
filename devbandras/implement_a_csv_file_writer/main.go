package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Flags struct {
	Source      string
	Destination string
	Zip         bool
}

type Records map[string]interface{}

func parseFlags() (*Flags, error) {
	sourceFlag := flag.String("s", "", "Path to the source JSON file")
	destinationFlag := flag.String("d", "", "Path to the destination CSV file")
	zipFlag := flag.Bool("z", false, "Archive the CSV into a ZIP file")
	flag.Parse()

	flags := Flags{
		Source:      *sourceFlag,
		Destination: *destinationFlag,
		Zip:         *zipFlag,
	}

	if flags.Source == "" || flags.Destination == "" {
		return nil, fmt.Errorf("source and destination flags are required")
	}

	return &flags, nil
}

func writeCSV(destinationFile *os.File, records Records) {
	// Write CSV file
}

func archiveCSV(destinationFile string) {
	// Archive CSV file
}

func main() {
	flags, err := parseFlags()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Usage: -s <source json_file> -d <destination csv_file> [-z]")
		os.Exit(1)
	}

	// Open and Read source JSON file
	sourceFile, err := os.Open(flags.Source)
	if err != nil {
		fmt.Printf("Error opening source file: %v\n", err)
		os.Exit(1)
	}
	defer sourceFile.Close()

	// Decode JSON file
	var records Records
	err = json.NewDecoder(sourceFile).Decode(&records)
	if err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		os.Exit(1)
	}

	// Create destination CSV file
	destinationFile, err := os.Create(flags.Destination)
	if err != nil {
		fmt.Printf("Error creating destination file: %v\n", err)
		os.Exit(1)
	}
	defer destinationFile.Close()

	// Write CSV file
	writeCSV(destinationFile, records)

	// Archive CSV file
	if flags.Zip {
		archiveCSV(flags.Destination)
	}

}
