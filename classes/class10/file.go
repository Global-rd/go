package main

import (
	"fmt"
	"io"
	"os"
)

func WriteFile(filename string) (err error) {
	var file *os.File
	_, err = os.Stat(filename)
	ok := os.IsNotExist(err)

	if !ok {
		file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	} else {
		file, err = os.Create(filename)
	}
	defer file.Close()
	if err != nil {
		return err
	}

	_, err = file.Write([]byte("hello! World!2"))
	if err != nil {
		return err
	}

	return
}

func CopyFile(old, new string) (err error) {
	var oldFile *os.File
	if _, err = os.Stat(old); os.IsNotExist(err) {
		return err
	}
	oldFile, err = os.Open(old)
	if err != nil {
		return err
	}
	defer oldFile.Close()

	if _, err = os.Stat(new); !os.IsNotExist(err) {
		return fmt.Errorf("given file %s already exists", new)
	}

	var newFile *os.File
	newFile, err = os.Create(new)
	if err != nil {
		return err
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, oldFile)
	return err
}
