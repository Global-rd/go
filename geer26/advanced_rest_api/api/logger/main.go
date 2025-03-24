package logger

import (
	"fmt"
	"sync"
	"time"
)

type Log struct {
	logchan  chan string
	Filelock sync.Mutex
}

func (l Log) INFO(info string) {
	l.logchan <- fmt.Sprintf("INFO :: %s", info)
}

func (l Log) ERROR(info string) {
	l.logchan <- fmt.Sprintf("ERROR :: %s", info)
}

func (l Log) WriteLog() {
	for log := range l.logchan {
		now := time.Now().String()
		fmt.Printf("\n%s :: %s", now, log)
	}
}

func (l Log) CloseLog() {
	close(l.logchan)
}

func InitLogger() (*Log, error) {
	logger := Log{
		logchan: make(chan string, 100),
	}

	go logger.WriteLog()

	return &logger, nil
}
