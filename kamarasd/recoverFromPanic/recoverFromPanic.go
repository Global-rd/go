package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

type Writer interface {
	Write(string) error
	Flush() error
	Close() error
}

type BatchWriter struct {
	buffer    *strings.Builder
	target    io.Writer
	batchSize int
}

func NewBatchWriter(target io.Writer, batchSize int) *BatchWriter {
	return &BatchWriter{
		buffer:    &strings.Builder{},
		target:    target,
		batchSize: batchSize,
	}
}

func (bw *BatchWriter) Write(data string) error {
	bw.buffer.WriteString(data)
	if bw.buffer.Len() >= bw.batchSize {
		return bw.Flush()
	}
	return nil
}

func (bw *BatchWriter) Flush() error {
	_, err := bw.target.Write([]byte(bw.buffer.String()))
	bw.buffer.Reset()
	return err
}

func (bw *BatchWriter) Close() error {
	err := bw.Flush()
	if closer, ok := bw.target.(io.Closer); ok {
		return closer.Close()
	}
	return err
}

type ChuckJoke struct {
	Created_at string `json:"created_at"`
	Value      string `json:"value"`
}

const (
	fileName  = "chuckJoke.txt"
	jokeUrl   = "https://api.chucknorris.io/jokes/random/"
	loopLimit = 10
)

func main() {
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Error opening file")
	}

	file.Close()
	file, err = os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		fmt.Println("Error opening file")
	}

	batchWriter := NewBatchWriter(file, 1024)

	writeFile(batchWriter)
}

func writeFile(writer Writer) {
	defer func() {
		a := recover()
		if a != nil {
			fmt.Println("Recovered from panic, Flush and close writer")
		}
		writer.Flush()
		writer.Close()
	}()

	var joke ChuckJoke
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < loopLimit; i++ {
		random := rand.Intn(loopLimit)
		if i != random {
			resp, err := http.Get(jokeUrl)
			if err != nil {
				fmt.Println("Error getting joke")
				return
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error reading response body")
				return
			}
			err = json.Unmarshal(body, &joke)
			if err != nil {
				fmt.Println("Error occured")
			}

			writer.Write(joke.Value)
			writer.Write("\n")
			resp.Body.Close()
		} else {
			panic("OMG I am panicking!!!!")
		}
	}
	fmt.Println("If you see this line, it means i'm not panicking")
	writer.Flush()
	writer.Close()

}
