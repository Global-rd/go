package main

import (
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

}
