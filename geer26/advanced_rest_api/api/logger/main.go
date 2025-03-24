package logger

import (
	"fmt"
	"time"
)

type Log struct{}

func (l Log) INFO(info string) {
	now := time.Now().String()
	fmt.Printf("%s :: INFO: %s\n", now, info)
}

func (l Log) ERROR(info string) {
	now := time.Now().String()
	fmt.Printf("%s :: ERROR: %s\n", now, info)
}

func InitLogger() (*Log, error) {
	logger := Log{}
	return &logger, nil
}
