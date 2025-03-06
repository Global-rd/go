package provider

import "testing"

const fileName = "../test_input.json"

func TestNewJsonFileReader(t *testing.T) {
	dataSource := NewJsonFileReader(fileName)
	err := dataSource.CheckSource()
	if err != nil {
		t.Error(err)
	}
}

func TestFileGetData(t *testing.T) {
	dataSource := NewJsonFileReader(fileName)
	data, err := dataSource.GetData()
	expectedLen := 20
	if err != nil {
		t.Error(err)
	}
	if len(data) == 0 {
		t.Error("no data")
	}
	t.Log(len(data))
	if len(data) != expectedLen {
		t.Errorf("expected length %d, got %d", expectedLen, len(data))
	}
}
