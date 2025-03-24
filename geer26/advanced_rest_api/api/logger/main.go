package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Option interface {
	apply(*Log)
}

type OptionFunc func(*Log)

func (f OptionFunc) apply(logger *Log) {
	f(logger)
}

func WithLogfile(p string) OptionFunc {
	return OptionFunc(func(l *Log) {
		l.LogfilePath = p
	})
}

type Log struct {
	Logchan     chan string
	LogfilePath string
}

func (l Log) INFO(info string) {
	l.Logchan <- fmt.Sprintf("INFO :: %s", info)
}

func (l Log) ERROR(info string) {
	l.Logchan <- fmt.Sprintf("ERROR :: %s", info)
}

func (l Log) WriteLog() {
	for logentry := range l.Logchan {
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
		err := l.AppendLog([]byte(fmt.Sprintf("%s :: %s\n", time, logentry)))
		if err != nil {
			log.Println("Error at append to logfile: ", err)
		}
	}
}

func (l Log) AppendLog(data []byte) error {
	file, err := os.OpenFile(l.LogfilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func (l Log) CloseLog() {
	l.Logchan <- "INFO :: Logging service shut down"
	close(l.Logchan)
}

func InitLogger(options ...Option) (*Log, error) {
	logger := Log{
		Logchan:     make(chan string),
		LogfilePath: "server.log",
	}

	for _, option := range options {
		option.apply(&logger)
	}

	err := logger.CreateLogfile()
	if err != nil {
		return &logger, err
	}

	go logger.WriteLog()
	logger.INFO("Log system started")

	return &logger, nil
}

func (l *Log) CreateLogfile() error {
	if _, err := os.Lstat(l.LogfilePath); err == nil {
		return nil
	}

	file, err := os.Create(l.LogfilePath)
	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}
