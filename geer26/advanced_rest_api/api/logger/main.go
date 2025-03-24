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
		now := time.Now()
		time := fmt.Sprintf(
			"%d/%s/%d, %d:%d:%d.%d", now.Year(),
			now.Month().String(),
			now.Day(),
			now.Hour(),
			now.Minute(),
			now.Second(),
			now.Nanosecond(),
		)
		fmt.Printf("\n%s :: %s", time, log)
	}
}

func (l Log) CloseLog() {
	l.logchan <- "INFO :: Logging service shut down"
	close(l.logchan)
}

func InitLogger() (*Log, error) {
	logger := Log{
		logchan: make(chan string, 100),
	}

	go logger.WriteLog()
	logger.INFO("Log system started")

	return &logger, nil
}
