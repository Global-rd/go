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

func WrtieZip(filename string) error {
	usersfilename := fmt.Sprintf("%s_users.csv", strings.Split(filename, ".")[0])
	productsfilename := fmt.Sprintf("%s_products.csv", strings.Split(filename, ".")[0])

	zipfilename := fmt.Sprintf("%s.zip", strings.Split(filename, ".")[0])
	archive, err := os.Create(zipfilename)
	if err != nil {
		return fmt.Errorf("error at creating archive-> %s", err.Error())
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()

	usersfile, err := os.Open(usersfilename)
	if err != nil {
		return fmt.Errorf("error at opening users file-> %s", err.Error())
	}
	defer usersfile.Close()
	productsfile, err := os.Open(productsfilename)
	if err != nil {
		return fmt.Errorf("error at opening products file-> %s", err.Error())
	}
	defer productsfile.Close()

	zippedusers, err := zipWriter.Create(usersfilename)
	if err != nil {
		return fmt.Errorf("error at compress users-> %s", err.Error())
	}
	if _, err := io.Copy(zippedusers, usersfile); err != nil {
		return fmt.Errorf("error at copy compressed users-> %s", err.Error())
	}

	zippedproducts, err := zipWriter.Create(productsfilename)
	if err != nil {
		return fmt.Errorf("error at compress products-> %s", err.Error())
	}
	if _, err := io.Copy(zippedproducts, productsfile); err != nil {
		return fmt.Errorf("error at copy compressed products-> %s", err.Error())
	}

	if err := os.Remove(usersfilename); err != nil {
		return fmt.Errorf("error at deleting plain users-> %s", err.Error())
	}
	if err := os.Remove(productsfilename); err != nil {
		return fmt.Errorf("error at deleting plain products-> %s", err.Error())
	}

	return nil
}

func Flush(d Data, filename string) error {
	usersfilename := fmt.Sprintf("%s_users.csv", strings.Split(filename, ".")[0])
	productsfilename := fmt.Sprintf("%s_products.csv", strings.Split(filename, ".")[0])

	usersfile, err := os.Create(usersfilename)
	if err != nil {
		return fmt.Errorf("error at creating users file-> %s", err.Error())
	}
	defer usersfile.Close()
	usersfile.Write([]byte("Id,Name,Email,Age"))
	usersfile.Write([]byte("\n"))
	for _, v := range d.Users {
		to_save := fmt.Sprintf("%d,%s,%s,%d", v.Id, v.Name, v.Email, v.Age)
		usersfile.Write([]byte(to_save))
		usersfile.Write([]byte("\n"))
	}

	productsfile, err := os.Create(productsfilename)
	if err != nil {
		return fmt.Errorf("error at creating products file-> %s", err.Error())
	}
	defer productsfile.Close()
	productsfile.Write([]byte("Id,Name,Price,Stock"))
	productsfile.Write([]byte("\n"))

	for _, v := range d.Products {
		to_save := fmt.Sprintf("%d,%s,%f,%d", v.Id, v.Name, v.Price, v.Stock)
		productsfile.Write([]byte(to_save))
		productsfile.Write([]byte("\n"))
	}

	return nil
}

func Read(filename string) (*Data, error) {
	var data Data
	file, err := os.ReadFile(filename)
	if err != nil {
		return &data, fmt.Errorf("error at opening source file-> %s", err.Error())
	}
	if err = json.Unmarshal(file, &data); err != nil {
		return &data, fmt.Errorf("error at parsing source file-> %s", err.Error())
	}
	return &data, nil
}

func CheckParams(sourceFlagShort, sourceFlagLong, ouputFlagShort, outputeFlagLong string) error {
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
	if err := CheckParams(*sourceFlagShort, *sourceFlagLong, *ouputFlagShort, *outputeFlagLong); err != nil {
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

	dat, err := Read(sourceFileName)
	if err != nil {
		panic(err.Error())
	}

	if err := Flush(*dat, outputFileName); err != nil {
		panic(err.Error())
	}

	if zipped {
		if err := WrtieZip(outputFileName); err != nil {
			panic(err.Error())
		}
	}

}
