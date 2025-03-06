package provider

import (
	"csv-writer/model"
	"errors"
	"fmt"
	"io"
	"os"
)

type JsonFileReader struct {
	FileName string
}

func (j *JsonFileReader) CheckSource() error {
	info, err := os.Stat(j.FileName)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.Join(fmt.Errorf("file not found: %s", j.FileName), err)
		} else {
			return err
		}
	}
	if !info.Mode().IsRegular() {
		return fmt.Errorf("not a regular file: %s", j.FileName)
	}
	var file *os.File
	file, err = os.Open(j.FileName)
	if err != nil {
		return errors.Join(fmt.Errorf("file not readable: %s", j.FileName), err)
	}
	defer file.Close()
	return nil
}

func (j *JsonFileReader) GetData() ([]model.JsonData, error) {
	f, err := os.Open(j.FileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var fileContent []byte

	fileContent, err = io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var data []model.JsonData
	data, err = model.ParseJsonData(string(fileContent))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func NewJsonFileReader(fileName string) *JsonFileReader {
	return &JsonFileReader{FileName: fileName}
}
