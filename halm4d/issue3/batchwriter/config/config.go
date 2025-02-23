package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	outputFileName     string
	reCreateOutputFile bool
	maxBatchSize       int
	serverUrl          string
	panicRate          int
}

func NewConfig() Config {
	c := Config{}
	c.readConfig()
	return c
}

func (c Config) readConfig() {
	c.outputFileName = os.Getenv("OUTPUT_FILE_NAME")
	if c.outputFileName == "" {
		fmt.Println("OUTPUT_FILE_NAME environment variable not set, applying default value: output.txt")
		c.outputFileName = "output.txt"
	}

	reCreateOutputFile, err := strconv.ParseBool(os.Getenv("RECREATE_OUTPUT_FILE"))
	if err != nil {
		fmt.Println("Error parsing RECREATE_OUTPUT_FILE environment variable, applying default value:", err)
		reCreateOutputFile = true
	}
	c.reCreateOutputFile = reCreateOutputFile

	maxBatchSize, err := strconv.Atoi(os.Getenv("MAX_BATCH_SIZE"))
	if err != nil {
		fmt.Println("Error parsing MAX_BATCH_SIZE environment variable, applying default value:", err)
		maxBatchSize = 3
	}
	c.maxBatchSize = maxBatchSize

	c.serverUrl = os.Getenv("SERVER_URL")
	if c.serverUrl == "" {
		fmt.Println("SERVER_URL environment variable not set, applying default value: http://localhost:8080")
		c.serverUrl = "http://localhost:8080"
	}

	panicRate, err := strconv.Atoi(os.Getenv("PANIC_RATE"))
	if err != nil {
		fmt.Println("PANIC_RATE environment variable not set, applying default value: 10")
		panicRate = 10
	}
	c.panicRate = panicRate
}

func (c Config) OutputFileName() string {
	return c.outputFileName
}

func (c Config) ReCreateOutputFile() bool {
	return c.reCreateOutputFile
}

func (c Config) MaxBatchSize() int {
	return c.maxBatchSize
}

func (c Config) ServerUrl() string {
	return c.serverUrl
}

func (c Config) PanicRate() int {
	return c.panicRate
}
