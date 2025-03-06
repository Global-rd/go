package provider

import (
	"fmt"
	"testing"
)

const url = "http://localhost:11434/api/generate"

func TestNewOllamaProvider(t *testing.T) {
	dataSource := NewOllamaProvider(url)
	err := dataSource.CheckSource()
	if err != nil {
		t.Error(err)
	}
	data, err := dataSource.sendRequest("generate a json list of 20 elements, each containing a field firstName, lastNAme, userName, phone and email. return only the generated json nothing else")
	if err != nil {
		t.Error(err)
	}
	t.Log(data)
}

func TestGetData(t *testing.T) {
	dataSource := NewOllamaProvider(url)
	data, err := dataSource.GetData()
	if err != nil {
		t.Error(err)
	}
	if len(data) == 0 {
		t.Error("no data")
	}
	fmt.Println(len(data))
}
