package main

import (
	"csv-exporter/config"
	"csv-exporter/ollama"
	"csv-exporter/tracker"
	"csv-exporter/writer"
	"csv-exporter/zipper"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {

	conf := config.NewConfig().Parse()
	if conf.Help() {
		conf.PrintHelp()
		os.Exit(0)
	}

	ollamaClient := ollama.NewOllama(
		conf.OllamaUrl(),
		conf.OllamaModel(),
		conf.Verbose(),
	)

	fmt.Println("Pulling model...")
	err := ollamaClient.PullModel()
	if err != nil {
		log.Fatalln("Error:", err)
	}
	fmt.Println("Model pulled successfully.")

	fmt.Println("Generating response...")
	generate, err := ollamaClient.Generate("Generate movies, books, and games, include these fields: title, type, description, genre, and release date. Return the result in JSON format. Put everything into one JSON array. Only return the JSON, do not add any other text.")
	if err != nil {
		log.Fatalln("Error:", err)
	}
	fmt.Println("Response:", generate)

	var tracks []tracker.Track
	err = json.Unmarshal([]byte(generate), &tracks)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	if conf.Verbose() {
		fmt.Println("Unmarshalled tracks:", tracks)
	}

	outputFile, err := os.Create("output.csv")
	if err != nil {
		log.Fatalln("Error:", err)
	}
	defer outputFile.Close()

	csvConverter := writer.NewCSVConverter(outputFile)
	defer csvConverter.Close()

	convertedTracks := tracker.Convert(tracks)
	if conf.Verbose() {
		fmt.Println("Converted tracks:", convertedTracks)
	}
	err = csvConverter.Write(convertedTracks)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	err = csvConverter.Close()
	if err != nil {
		log.Fatalln("Error:", err)
	}

	fmt.Println("CSV file created successfully.")

	fmt.Println("Zipping CSV file...")
	zip := zipper.NewZipper(conf)
	err = zip.ZipFile("output.csv")
	if err != nil {
		log.Fatalln("Error:", err)
	}

	fmt.Println("CSV file zipped successfully.")
	fmt.Println("Done.")
}
