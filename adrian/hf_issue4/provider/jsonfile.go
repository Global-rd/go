package provider

import (
	"errors"
	"fmt"
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

func (j *JsonFileReader) GetData() (string, error) {
	return "", nil
}

func (j *JsonFileReader) Close() {

}

func NewJsonFileReader(fileName string) *JsonFileReader {
	return &JsonFileReader{FileName: fileName}
}
