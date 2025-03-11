package main

import (
	"archive/zip"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

const (
	inputDataFileName = "data.json"
	fileName          = "users"
	csvExtension      = ".csv"
	zipExtension      = ".zip"
)

func main() {

	err := createCsv(fileName + csvExtension)
	if err != nil {
		fmt.Println("Error creating and saving file:", err)
		os.Exit(1)
	}

	err = zipCsvFile(fileName+csvExtension, fileName+zipExtension)
	if err != nil {
		fmt.Println("Error zipping file:", err)
		os.Exit(1)
	}

	fmt.Println("File saved successfully")
	err = os.Remove(fileName + csvExtension)
	if err != nil {
		fmt.Println("Error deleting csv file:", err)
		os.Exit(1)
	}
}

func createCsv(csvFileName string) error {

	// open json file
	file, err := os.Open(inputDataFileName)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}

	defer file.Close()

	var userData []User

	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	err = json.Unmarshal(data, &userData)
	if err != nil {
		return fmt.Errorf("error unmarshalling json: %v", err)
	}

	// create CSV and write data
	csvFile, err := os.Create(csvFileName)
	if err != nil {
		return fmt.Errorf("error creating csv file: %v", err)
	}

	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	bomUTF8 := []byte{0xEF, 0xBB, 0xBF}
	csvWriter.Write([]string{string(bomUTF8[:]) + "Név", "Életkor", "Város"})

	for _, user := range userData {
		err = csvWriter.Write([]string{user.Name, fmt.Sprintf("%d", user.Age), user.City})
		if err != nil {
			return fmt.Errorf("error writing csv row: %v", err)
		}
	}

	csvWriter.Flush()
	csvFile.Close()

	return nil
}

func zipCsvFile(csvFileName string, zipFileName string) error {

	zipFile, err := os.Create(zipFileName)
	if err != nil {
		return fmt.Errorf("error creating zip file: %v", err)
	}
	defer zipFile.Close()
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	csvFileOpen, err := os.Open(csvFileName)
	if err != nil {
		return fmt.Errorf("error opening csv file: %v", err)
	}

	csvWriterActual, err := zipWriter.Create(csvFileName)
	if err != nil {
		return fmt.Errorf("error creating csv file in zip: %v", err)
	}

	_, err = io.Copy(csvWriterActual, csvFileOpen)
	if err != nil {
		return fmt.Errorf("error copying csv file to zip: %v", err)
	}

	csvFileOpen.Close()

	return err
}
