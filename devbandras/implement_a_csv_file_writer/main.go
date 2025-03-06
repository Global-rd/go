package main

import (
	"csv_writer/utils"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	flags := utils.NewFlags()
	flags.Parse()
	err := flags.Validate()

	if err != nil {
		fmt.Println(err)
		utils.PrintUsage()
		os.Exit(1)
	}

	// json fájl megnyitása
	sourceFile, err := os.Open(flags.Source)
	if err != nil {
		fmt.Printf("Error opening source file: %v\n", err)
		os.Exit(1)
	}
	defer sourceFile.Close()

	// json fájl dekódolása
	var records []utils.Records
	err = json.NewDecoder(sourceFile).Decode(&records)
	if err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		os.Exit(1)
	}

	// csv fájl létrehozása, írása
	csvWriter := utils.NewCsvWriter()
	csvWriter.FileName = flags.Destination
	csvWriter.Comma = ";"
	csvWriter.Records = records
	err = csvWriter.Write()
	if err != nil {
		fmt.Printf("Error writing CSV: %v\n", err)
		os.Exit(1)
	}

	// csv fál tömörítése
	if flags.Zip {
		zip := utils.NewZip()
		zip.CSVFileName = flags.Destination
		err = zip.Execute()
		if err != nil {
			fmt.Printf("Error zipping CSV: %v\n", err)
			os.Exit(1)
		}
	}

	// ha tömörítés sikeres volt, töröljük a csv fájlt
	err = os.Remove(flags.Destination)
	if err != nil {
		fmt.Printf("Error deleting CSV file: %v\n", err)
		os.Exit(1)
	}
}
