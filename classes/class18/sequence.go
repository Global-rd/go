package main

import (
	"errors"
	"fmt"
	"time"
)

type Signal struct{}

var Done = Signal{}

type Task struct {
	Action string
	Result chan string
	Error  chan error
}

func RunRestoreSequence() {
	ch := make(chan Task)

	result := make(chan string)
	errCh := make(chan error)

	task := Task{
		Action: "find the latest payment",
		Result: result,
		Error:  errCh,
	}

	worker(ch)

	ch <- task

	time.Sleep(time.Second * 5)

	select {
	case v := <-task.Result:
		fmt.Println("work is done", v)
	case err := <-task.Error:
		fmt.Println("error", err)
	}

	fmt.Println("finished")
}

func worker(ch <-chan Task) {
	go func() {
		task := <-ch

		if len(task.Action) == 0 {
			task.Error <- errors.New("there is no action to work on")
			return
		}

		task.Result <- "payment id: 511"
	}()
}
