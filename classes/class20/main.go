package main

import (
	"errors"
	"fmt"
	"time"
)

type Signal struct{}

var Done = Signal{}

type Task struct {
	Request string
	done    chan Signal
	err     chan error
}

func (t Task) Done() <-chan Signal {
	return t.done
}

func (t Task) Error() <-chan error {
	return t.err
}

func NewTask(request string) Task {
	done := make(chan Signal)
	err := make(chan error)

	return Task{
		Request: request,
		done:    done,
		err:     err,
	}
}

func main() {
	// task := NewTask("do the action")

	ch, done := GenerateTask()

	go func() {
		for receivedTask := range ch {
			if len(receivedTask.Request) == 0 {
				select {
				case <-time.After(time.Second * 3):
					fmt.Println("worker", "timed out")
				case receivedTask.err <- errors.New("no task to execute"):
				}
				return
			}
			select {
			case receivedTask.done <- Done:
			case <-time.After(time.Second * 3):
				fmt.Println("worker", "timed out")
			}
		}
	}()

	<-done

	fmt.Println("finished")
}

func GenerateTask() (<-chan Task, <-chan Signal) {
	ch := make(chan Task)
	done := make(chan Signal)

	go func() {
		for i := range 10 {
			task := NewTask(fmt.Sprintf("action: %d", i))

			ch <- task

			select {
			case <-task.Done():
				fmt.Println("task is finished", i)
			case err := <-task.Error():
				fmt.Println("error", err.Error())
			case <-time.After(time.Second * 3):
				fmt.Println("error", "timed out")
			}
		}
		close(ch)
		done <- Done
	}()

	return ch, done
}
