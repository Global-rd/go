package main

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type User struct {
	Id    int    `json:"id"`
	Age   int    `json:"age"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Products struct {
	Id    int     `json:"id"`
	Stock int     `json:"stock"`
	Price float64 `json:"price"`
	Name  string  `json:"name"`
}

type Data struct {
	Users    []User     `json:"users"`
	Products []Products `json:"products"`
}

func flush(d Data, filename string, zipped bool) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write([]byte("USERS"))
	file.Write([]byte("\n"))
	file.Write([]byte("Id,Name,Email,Age"))
	file.Write([]byte("\n"))
	for _, v := range d.Users {
		to_save := fmt.Sprintf("%d,%s,%s,%d", v.Id, v.Name, v.Email, v.Age)
		file.Write([]byte(to_save))
		file.Write([]byte("\n"))
	}
	file.Write([]byte("USERS"))
	file.Write([]byte("\n"))
	file.Write([]byte("Id,Name,Price,Stock"))
	file.Write([]byte("\n"))
	for _, v := range d.Products {
		to_save := fmt.Sprintf("%d,%s,%f,%d", v.Id, v.Name, v.Price, v.Stock)
		file.Write([]byte(to_save))
		file.Write([]byte("\n"))
	}

	if zipped {
		name := strings.TrimSpace(strings.Split(filename, ".")[0])
		archive, err := os.Create(fmt.Sprintf("%s.zip", name))
		if err != nil {
			return err
		}
		defer archive.Close()
		zipWriter := zip.NewWriter(archive)
		defer zipWriter.Close()
		file, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer file.Close()
		zipfile, err := zipWriter.Create(filename)
		if err != nil {
			return err
		}
		if _, err := io.Copy(zipfile, file); err != nil {
			return err
		}
		if err := os.Remove(filename); err != nil {
			return err
		}
	}
	return nil
}

func read(filename string) (*Data, error) {
	var data Data
	file, err := os.ReadFile(filename)
	if err != nil {
		return &data, err
	}
	if err = json.Unmarshal(file, &data); err != nil {
		return &data, err
	}
	return &data, nil
}

func checkParams(sourceFlagShort, sourceFlagLong, ouputFlagShort, outputeFlagLong string) error {
	if sourceFlagShort == "" && sourceFlagLong == "" && ouputFlagShort == "" && outputeFlagLong == "" {
		return errors.New("source and output file must be specified")
	}
	if sourceFlagShort == "" && sourceFlagLong == "" {
		return errors.New("source file must be specified")
	}
	if sourceFlagShort != "" && sourceFlagLong != "" {
		return errors.New("too many arguments for source file")
	}
	if ouputFlagShort == "" && outputeFlagLong == "" {
		return errors.New("output file must be specified")
	}
	if ouputFlagShort != "" && outputeFlagLong != "" {
		return errors.New("too many arguments for output file")
	}
	return nil
}

func main() {
	sourceFlagShort := flag.String("s", "", "Input file path")
	sourceFlagLong := flag.String("source", "", "Input file path")
	ouputFlagShort := flag.String("o", "", "Output file path")
	outputeFlagLong := flag.String("output", "", "Output file path")
	zippedShort := flag.Bool("z", false, "Zip the output file")
	zippedLong := flag.Bool("zip", false, "Zip the output file")
	flag.Parse()
	if err := checkParams(*sourceFlagShort, *sourceFlagLong, *ouputFlagShort, *outputeFlagLong); err != nil {
		panic(err.Error())
	}
	var sourceFileName string
	if *sourceFlagShort != "" {
		sourceFileName = *sourceFlagShort
	} else {
		sourceFileName = *sourceFlagLong
	}
	var outputFileName string
	if *ouputFlagShort != "" {
		outputFileName = *ouputFlagShort
	} else {
		outputFileName = *outputeFlagLong
	}
	var zipped bool
	if *zippedLong || *zippedShort {
		zipped = true
	}

	dat, err := read(sourceFileName)
	if err != nil {
		panic(err.Error())
	}
	flush(*dat, outputFileName, zipped)
}
