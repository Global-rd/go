package provider

import (
	"csv-writer/model"
	"errors"
	"strings"
)

type DataProvider interface {
	GetData() ([]model.JsonData, error)
	CheckSource() error
}

func NewProvider(source string) (DataProvider, error) {
	if source == "" {
		return nil, errors.New("no source provided")
	}

	if strings.HasPrefix(source, "http") {
		return NewOllamaProvider(source), nil
	} else {
		return NewJsonFileReader(source), nil
	}
}
