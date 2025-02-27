package main

import (
	"archive/zip"
	"io"
	"os"
	"time"
)

func createZip(filename string) error {
	zipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	header := &zip.FileHeader{
		Name:     "apple.txt",
		Method:   zip.Deflate,
		Modified: time.Now(),
	}

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = writer.Write([]byte("file content"))

	return err
}

func toZip(zipFilename string, files []string) error {
	zipFile, err := os.Create(zipFilename)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, fileName := range files {
		if _, err := os.Stat(fileName); os.IsNotExist(err) {
			return err
		}
		file, err := os.Open(fileName)
		if err != nil {
			return err
		}
		defer file.Close()

		header := &zip.FileHeader{
			Name:     fileName,
			Method:   zip.Deflate,
			Modified: time.Now(),
		}

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		_, err = io.Copy(writer, file)
		if err != nil {
			return err
		}

	}
	return nil
}
