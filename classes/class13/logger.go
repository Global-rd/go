package main

import "errors"

type LogType int

const (
	JSONLoggerType LogType = iota
	RawLoggerType
)

type JSONLogger struct {
}

func (j JSONLogger) Log(message string) {
	// .. code
}

type RawLogger struct {
}

func (j RawLogger) Log(message string) {
	// .. code
}

type Logger interface {
	Log(message string)
}

func newJSONLogger() JSONLogger {
	return JSONLogger{}
}

func NewRawLogger() RawLogger {
	return RawLogger{}
}

func CreateLogger(loggerType LogType) (Logger, error) {
	switch loggerType {
	case JSONLoggerType:
		return newJSONLogger(), nil
	case RawLoggerType:
		return NewRawLogger(), nil
	}
	return nil, errors.New("unsupported logger type")
}
