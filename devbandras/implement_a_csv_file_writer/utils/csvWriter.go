package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Records map[string]any

type Writer interface {
	Write() error
}

type CSVWriter struct {
	Records  []Records
	FileName string
	Comma    string
	Writer   *csv.Writer
}

func NewCsvWriter() *CSVWriter {
	return &CSVWriter{}
}

// A Write függvény létrehoz egy új CSV fájlt a megadott rekordokból.
// A függvény ellenőrzi, hogy a fájl már létezik-e, és ha nem, akkor létrehozza, és feltölti a c.Records tartalmával.
//
// Paraméterek:
// - c: a CSVWriter struktúra mutatója
//
// Returns:
// - error: a CSV fájl írása során keletkező hibát tartalmazza.

func (c *CSVWriter) Write() error {
	// Check if there are any records to write
	if len(c.Records) == 0 {
		return fmt.Errorf("no records to write")
	}

	// Check if the file already exists
	if _, err := os.Stat(c.FileName); err == nil {
		return fmt.Errorf("file already exists: %s", c.FileName)
	}

	// Create a new file
	file, err := os.Create(c.FileName)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	c.Writer = csv.NewWriter(file)
	defer c.Writer.Flush()

	// Extract headers from the first record
	var headers []string
	for key := range c.Records[0] {
		headers = append(headers, key)
	}

	// Write CSV headers
	if err := c.Writer.Write(headers); err != nil {
		return err
	}

	// Write CSV rows
	for _, record := range c.Records {
		var row []string
		for _, header := range headers {
			row = append(row, fmt.Sprintf("%v", record[header]))
		}
		if err := c.Writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}
