package provider

import "testing"

func TestNewJsonFileReader(t *testing.T) {
	fileName := "../test_input.json"
	dataSource := NewJsonFileReader(fileName)
	err := dataSource.CheckSource()
	if err != nil {
		t.Error(err)
	}
}
