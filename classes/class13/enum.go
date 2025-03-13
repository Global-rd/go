package main

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warning
	Error
	Fatal
)

func (l LogLevel) String() string {
	levels := []string{"Debug", "Info", "Warning", "Error", "Fatal"}

	return levels[l]
}
