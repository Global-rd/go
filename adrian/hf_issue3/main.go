package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

const (
	QuoteUrl       = "http://kollarovics.net:54380/quote"
	OutputFileName = "quotes_output.txt"
	QuotesToGet    = 50
	FlushInterval  = 5
	RndMin         = 3
	RndMax         = 5
)

type Writer interface {
	Write(string) error
	Flush() error
	Close() error
	GetSuccessCount() uint
}

type BatchWriter struct {
	WriteSuccess uint
	Buffer       []Quote
	OutputFile   os.File
}

func (bw *BatchWriter) Write(text string) error {
	quoteParts := strings.Split(text, "~")
	if len(quoteParts) != 2 {
		return errors.New("Invalid quote format")
	}
	quote := Quote{
		Text:   strings.TrimRight(quoteParts[0], " "),
		Author: quoteParts[1],
	}
	bw.Buffer = append(bw.Buffer, quote)
	return nil
}

func NewBatchWriter() (*BatchWriter, error) {
	output, err := os.OpenFile(OutputFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return nil, err
	}
	batchWriter := BatchWriter{
		WriteSuccess: 0,
		OutputFile:   *output,
	}
	batchWriter.Buffer = make([]Quote, 0)
	return &batchWriter, nil
}

func (bw *BatchWriter) Flush() error {
	for _, quote := range bw.Buffer {
		_, err := bw.OutputFile.WriteString(fmt.Sprintf("Author: %s, Quote: \"%s\"\n", quote.Author, quote.Text))
		if err != nil {
			return err
		}
		bw.WriteSuccess++
	}
	bw.Buffer = make([]Quote, 0)
	return nil
}

func (bw *BatchWriter) Close() error {
	return bw.OutputFile.Close()
}

func (bw *BatchWriter) GetSuccessCount() uint {
	return bw.WriteSuccess
}

type Quote struct {
	Text   string
	Author string
}

func getQuote() (quote string, err error) {
	var resp *http.Response
	resp, err = http.Get(QuoteUrl)
	if err != nil || resp.StatusCode != http.StatusOK {
		return "", errors.New("Quote fetch failed")
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			fmt.Printf("Error closing response body: %v", err)
		}
	}()
	var body []byte
	body, err = io.ReadAll(resp.Body)
	quote = string(body)
	return quote, err
}

func shouldSimulatePanic(idx int) bool {
	// Checks randomly if the bits 3-5 are set in the loop index
	// if yes, a panic should be simulated
	bitIdx := rand.Intn(RndMax-RndMin+1) + RndMin
	checkBit := 1 << bitIdx
	return (checkBit & idx) != 0
}

func fetchAndWriteQuotes(writer Writer) (err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic occured: %v\n", r)
			fmt.Printf("So far %d quotes written. Flushing remaining quotes from bufer\n", writer.GetSuccessCount())
			err = writer.Flush()
		}
	}()
	for i := 0; i < QuotesToGet; i++ {
		var quote string
		quote, err = getQuote()
		err = writer.Write(quote)
		if err != nil {
			fmt.Printf("Error writing quote: %v\n", err)
			return err
		}
		if shouldSimulatePanic(i) {
			panic("Simulated panic")
		}
		if (i+1)%FlushInterval == 0 {
			err = writer.Flush()
			if err != nil {
				fmt.Printf("Error flushing quotes: %v\n", err)
				return err
			}
			fmt.Printf("So far %d quotes written\n", writer.GetSuccessCount())
		}
	}
	err = writer.Flush()
	if err != nil {
		fmt.Printf("Error flushing quotes: %v\n", err)
		return err
	}
	return err
}

func main() {
	fmt.Printf("Application starting, trying to get %d quotes\n", QuotesToGet)
	batchWriter, err := NewBatchWriter()
	if err != nil {
		fmt.Printf("Error initializing batch writer: %v\n", err)
		os.Exit(1)
	}

	defer func() {
		closeErr := batchWriter.Close()
		if closeErr != nil {
			fmt.Printf("Error closing output file: %v\n", closeErr)
		}
	}()
	err = fetchAndWriteQuotes(batchWriter)
	if err != nil {
		fmt.Printf("Error occured while fetching quotes: %v\nWritten quotes: %d\n,", err, batchWriter.WriteSuccess)
		os.Exit(1)
	}
	fmt.Printf("Application finished, wrote %d quotes\n", batchWriter.WriteSuccess)
}
