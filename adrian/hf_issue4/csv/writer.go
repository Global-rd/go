package csv

import (
	"archive/zip"
	"csv-writer/model"
	"csv-writer/provider"
	"fmt"
	"io"
	"os"
	"time"
)

const csvHeader = "First Name, Last Name, User Name, Phone Number, Email Address\n"

type DataWriter struct {
	zipped     bool
	dataSource provider.DataProvider
}

func NewDataWriter(dataSource provider.DataProvider, zipped bool) *DataWriter {
	return &DataWriter{
		zipped:     zipped,
		dataSource: dataSource,
	}
}

func (dw *DataWriter) writeData(writer io.Writer) error {
	_, err := writer.Write([]byte(csvHeader))
	if err != nil {
		return err
	}
	var data []model.JsonData
	data, err = dw.dataSource.GetData()
	if err != nil {
		return err
	}

	for _, rowData := range data {
		lineToWrite := fmt.Sprintf("%s, %s, %s, %s, %s\n", rowData.FirstName, rowData.LastName, rowData.UserName, rowData.PhoneNumber, rowData.EmailAddress)
		_, err = writer.Write([]byte(lineToWrite))
		if err != nil {
			return err
		}
	}
	return nil
}

func (dw *DataWriter) writeZip(outputName string) error {
	zipName := outputName + ".zip"
	zipFile, err := os.Create(zipName)
	if err != nil {
		return err
	}
	defer zipFile.Close()
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()
	var fileWriter io.Writer
	header := &zip.FileHeader{
		Name:     outputName,
		Method:   zip.Deflate,
		Modified: time.Now(),
	}
	fileWriter, err = zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	err = dw.writeData(fileWriter)
	if err != nil {
		return err
	}
	return nil
}

func (dw *DataWriter) writePlainCsv(outputName string) error {
	file, err := os.Create(outputName)
	if err != nil {
		return err
	}
	defer file.Close()
	err = dw.writeData(file)
	if err != nil {
		return err
	}
	return nil
}

func (dw *DataWriter) Write(outputName string) error {
	if dw.zipped {
		return dw.writeZip(outputName)
	} else {
		return dw.writePlainCsv(outputName)
	}
}
