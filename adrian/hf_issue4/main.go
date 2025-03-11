package main

import (
	"csv-writer/csv"
	"csv-writer/provider"
	"flag"
	"fmt"
	"os"
)

func parseFlags() (string, string, bool) {
	var zipped bool
	var source string
	var output string
	flag.BoolVar(&zipped, "z", false, "Zip the output file")
	flag.BoolVar(&zipped, "zip", false, "Zip the output file")
	flag.StringVar(&source, "s", "", "Source file")
	flag.StringVar(&source, "source", "", "Source file")
	flag.StringVar(&output, "o", "", "Output file")
	flag.StringVar(&output, "output", "", "Output file")
	flag.Parse()
	return source, output, zipped
}

func main() {
	source, output, zipped := parseFlags()
	if output == "" {
		fmt.Println("output file name must be specified")
		os.Exit(1)
	}

	dataProvider, err := provider.NewProvider(source)
	if err != nil {
		fmt.Printf("failed to initialize data source: %s\n", err)
		os.Exit(1)
	}
	err = dataProvider.CheckSource()
	if err != nil {
		fmt.Printf("error while check data source: %s\n", err)
		os.Exit(1)
	}

	writer := csv.NewDataWriter(dataProvider, zipped)
	err = writer.Write(output)
	if err != nil {
		fmt.Printf("error while writing data: %s\n", err)
	}

}
