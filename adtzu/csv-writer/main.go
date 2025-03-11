package main

import (
    "archive/zip"
    "encoding/csv"
    "encoding/json"
    "flag"
    "fmt"
    "io"
    "os"
)

type Record struct {
    Name   string `json:"name"`
    Age    int    `json:"age"`
    City   string `json:"city"`
}

func main() {
    sourceFlag := flag.String("s", "", "Source JSON file path")
    outputFlag := flag.String("o", "", "Output CSV file path")
    zipFlag := flag.Bool("z", false, "Archive CSV into ZIP")
    flag.StringVar(sourceFlag, "source", "", "Source JSON file path")
    flag.StringVar(outputFlag, "output", "", "Output CSV file path")
    flag.BoolVar(zipFlag, "zip", false, "Archive CSV into ZIP")
    flag.Parse()

    if *sourceFlag == "" || *outputFlag == "" {
        fmt.Println("Source and output flags are required")
        flag.PrintDefaults()
        os.Exit(1)
    }

    jsonFile, err := os.Open(*sourceFlag)
    if err != nil {
        fmt.Println("Error opening JSON file:", err)
        os.Exit(1)
    }
    defer jsonFile.Close()

    var records []Record
    err = json.NewDecoder(jsonFile).Decode(&records)
    if err != nil {
        fmt.Println("Error decoding JSON:", err)
        os.Exit(1)
    }

    var output io.Writer
    if *zipFlag {
        zipFile, err := os.Create(*outputFlag + ".zip")
        if err != nil {
            fmt.Println("Error creating ZIP file:", err)
            os.Exit(1)
        }
        defer zipFile.Close()

        zipWriter := zip.NewWriter(zipFile)
        defer zipWriter.Close()

        csvFile, err := zipWriter.Create(*outputFlag)
        if err != nil {
            fmt.Println("Error creating CSV in ZIP:", err)
            os.Exit(1)
        }
        output = csvFile
    } else {
        csvFile, err := os.Create(*outputFlag)
        if err != nil {
            fmt.Println("Error creating CSV file:", err)
            os.Exit(1)
        }
        defer csvFile.Close()
        output = csvFile
    }

    csvWriter := csv.NewWriter(output)
    defer csvWriter.Flush()

    // Write header
    err = csvWriter.Write([]string{"Name", "Age", "City"})
    if err != nil {
        fmt.Println("Error writing CSV header:", err)
        os.Exit(1)
    }

    // Write records
    for _, record := range records {
        err := csvWriter.Write([]string{record.Name, fmt.Sprintf("%d", record.Age), record.City})
        if err != nil {
            fmt.Println("Error writing CSV record:", err)
            os.Exit(1)
        }
    }

    fmt.Println("CSV file created successfully")
}
