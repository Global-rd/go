package main

import "fmt"

type BatchWriter struct {
	buffer []string
}

func (b BatchWriter) Write(v string) error {
	b.buffer = append(b.buffer, v)
	return nil
}

func (b BatchWriter) Flush() error {
	for _, value := range b.buffer {
		fmt.Println(value)
	}

	return nil
}
