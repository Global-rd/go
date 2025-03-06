package main

import (
	"csv-exporter/config"
	"csv-exporter/ollama"
	"csv-exporter/tracker"
	"csv-exporter/writer"
	"csv-exporter/zipper"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	// Parse config
	conf := config.NewConfig().Parse()
	if conf.Help() {
		conf.PrintHelp()
		os.Exit(0)
	}

	// Get JSON source
	jsonSource, err := getJsonSource(conf)
	if err != nil {
		log.Fatalln("Error while getting JSON source:", err)
	}

	// Unmarshal JSON
	var tracks []tracker.Track
	err = json.Unmarshal([]byte(jsonSource), &tracks)
	if err != nil {
		log.Fatalln("Error while unmarshalling JSON:", err)
	}
	if conf.Verbose() {
		fmt.Println("Unmarshalled tracks:", tracks)
	}

	// Create output file
	outputFile, err := os.Create(fmt.Sprintf("%s.csv", conf.Output()))
	if err != nil {
		log.Fatalln("Error while creating output file:", err)
	}
	defer func() {
		if err = outputFile.Close(); err != nil {
			log.Fatalln("Error while closing output file:", err)
		}
	}()

	// Write CSV file
	err = writeCSVFile(tracks, outputFile, conf)
	if err != nil {
		log.Fatalln("Error while writing CSV file:", err)
	}
	fmt.Println("CSV file created successfully.")

	// Zip CSV file if required
	if conf.Zip() {
		fmt.Println("Zipping CSV file...")
		zip := zipper.NewZipper(conf)
		err = zip.ZipFile("output.csv")
		if err != nil {
			log.Fatalln("Error while zipping CSV file:", err)
		}
		fmt.Println("CSV file zipped successfully.")
	}
	fmt.Println("Done.")
}

func writeCSVFile(tracks []tracker.Track, outputFile io.Writer, conf *config.Config) (err error) {
	csvWriter := writer.NewCSVWriter(outputFile)
	defer func() {
		if closeErr := csvWriter.Close(); closeErr != nil {
			err = errors.Join(err, closeErr)
		}
	}()

	// Convert tracks to CSV format
	convertedTracks := tracker.Convert(tracks)
	if conf.Verbose() {
		fmt.Println("Converted tracks:", convertedTracks)
	}

	// Write header
	header := []string{"Title", "Type", "Description", "Genre", "Release Date"}
	err = csvWriter.Write([][]string{header})
	if err != nil {
		return errors.New("Error writing CSV header: " + err.Error())
	}

	// Write tracks data
	err = csvWriter.Write(convertedTracks)
	if err != nil {
		return errors.New("Error writing CSV file: " + err.Error())
	}

	// Close CSV writer
	err = csvWriter.Close()
	if err != nil {
		return errors.New("Error closing CSV file: " + err.Error())
	}
	return nil
}

func getJsonSource(conf *config.Config) (string, error) {
	if conf.SourceType() == "ollama" {
		ollamaClient := ollama.NewOllama(
			conf.OllamaUrl(),
			conf.OllamaModel(),
			conf.Verbose(),
		)

		fmt.Println("Pulling model...")
		err := ollamaClient.PullModel()
		if err != nil {
			return "", errors.New("Error pulling model: " + err.Error())
		}
		fmt.Println("Model pulled successfully.")

		fmt.Println("Generating response...")
		generate, err := ollamaClient.Generate("Generate movies, books, and games, include these fields: title, type, description, genre, and release date. Return the result in JSON format. Put everything into one JSON array. Only return the JSON, do not add any other text.")
		if err != nil {
			return "", errors.New("Error generating response: " + err.Error())
		}
		fmt.Println("Response generated successfully.")
		if conf.Verbose() {
			fmt.Println("Response:", generate)
		}
		return generate, nil
	} else if conf.SourceType() == "file" {
		file, err := os.ReadFile(conf.SourceType())
		if err != nil {
			return "", errors.New("Error reading file: " + err.Error())
		}
		if conf.Verbose() {
			fmt.Println("File content:", string(file))
		}
		return string(file), nil
	} else {
		return "", errors.New("invalid source type. Use 'ollama' or 'file'")
	}
}
