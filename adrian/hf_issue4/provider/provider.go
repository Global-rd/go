package provider

import (
	"errors"
	"strings"
)

type DataProvider interface {
	GetData() (string, error)
	CheckSource() error
	Close()
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
