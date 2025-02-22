package main

import (
	"encoding/json"
	"errors"
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

	if err != nil {
		fmt.Println("Error opening file")
	}

	batchWriter := NewBatchWriter(file, 1024)

	err = writeFile(batchWriter)

	if err != nil {
		fmt.Println("An error occured: ", err)
	}

}

func writeFile(writer Writer) (err error) {
	defer func() {
		a := recover()
		if a != nil {
			err = errors.New("recovered from panic, flush and close writer")
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
				return errors.New("error getting joke")
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return errors.New("error reading response body")
			}
			err = json.Unmarshal(body, &joke)
			if err != nil {
				return errors.New("unmarshalling json run into an error")
			}

			writer.Write(joke.Value)
			writer.Write("\n")
			resp.Body.Close()
		} else {
			panic("OMG I am panicking!!!!")
		}
	}
	fmt.Println("If you see this line, it means i'm not panicking")
	err = writer.Flush()
	if err != nil {
		return errors.New("error happened when flush data")
	}
	err = writer.Close()
	if err != nil {
		return errors.New("error happened when close the file")
	}

	if err != nil {
		return err
	}
	return nil
}
