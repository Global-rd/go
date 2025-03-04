package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

type Flags struct {
	Source      string
	Destination string
	Zip         bool
}

type Records map[string]interface{}

// A parseFlags elemzi a parancssori kapcsolókat, és egy Flags struktúrára tér vissza, amely tartalmazza a feldolgozott kapcsolókat.
// Ellenőrzi a forrás- és célkapcsolókat, és hibát ad vissza, ha nincsenek megadva.
//
// Parameters:
// - None
//
// Return:
// - Flags: A feldolgozott kapcsolókat tartalmazó struktúra.
// - error: Egy hiba, ha a forrás- és célkapcsolók nincsenek megadva.
func parseFlags() (Flags, error) {
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
		return flags, fmt.Errorf("source and destination file are required")
	}

	return flags, nil
}

// A writeCSV a megadott rekordokat írja a megadott íróba CSV formátumban.
// Először kinyeri a fejléceket az első rekordból, és ezeket írja a CSV fájlba.
// Ezután az adatsorokat írja a kinyert fejlécek segítségével.
// Parameters:
// - w:  io.writer
// - records: A feldolgozandó rekordok
//
// Return: nil: Ha a rekordok sikeresen ki lettek írva a CSV-be; error: Ha hiba történt a rekordok kiírása közben.
func writeCSV(w io.Writer, records []Records) error {
	// van-e rekord
	if len(records) == 0 {
		return fmt.Errorf("no records to write")
	}

	// fejléc kinyerése
	var headers []string
	for key := range records[0] {
		headers = append(headers, key)
	}

	csvWriter := csv.NewWriter(w)
	defer csvWriter.Flush()

	// csv fejlec írása
	if err := csvWriter.Write(headers); err != nil {
		return err
	}

	// csv sorok írása
	for _, record := range records {
		var row []string
		for _, key := range headers {
			row = append(row, fmt.Sprintf("%v", record[key]))
		}
		if err := csvWriter.Write(row); err != nil {
			return err
		}
	}

	return nil
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

	// json fájl megnyitása
	sourceFile, err := os.Open(flags.Source)
	if err != nil {
		fmt.Printf("Error opening source file: %v\n", err)
		os.Exit(1)
	}
	defer sourceFile.Close()

	// json fájl dekódolása
	var records []Records
	err = json.NewDecoder(sourceFile).Decode(&records)
	if err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		os.Exit(1)
	}

	// csv fájl ellenőrzése
	if _, err := os.Stat(flags.Destination); err == nil {
		fmt.Printf("Destination file already exists: %s\n", flags.Destination)
		os.Exit(1)
	}

	// csv fájl létrehozása
	destinationFile, err := os.Create(flags.Destination)
	if err != nil {
		fmt.Printf("Error creating destination file: %v\n", err)
		os.Exit(1)
	}
	defer destinationFile.Close()

	// csv fájl írása
	err = writeCSV(destinationFile, records)
	if err != nil {
		fmt.Printf("Error writing CSV: %v\n", err)
		os.Exit(1)
	}

	// csv fál tömörítése
	if flags.Zip {
		archiveCSV(flags.Destination)
	}

}
