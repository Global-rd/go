package logger

import (
	"archive/zip"
	"fmt"
	"io"
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

func WithLogSize(s int) OptionFunc {
	return OptionFunc(func(l *Log) {
		if s > 1 {
			l.MaxLogfileSize = s
		}
	})
}

type Log struct {
	Logchan        chan string
	LogfilePath    string
	MaxLogfileSize int
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
			fmt.Println("Error at append to logfile: ", err)
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

func (l *Log) CreateLogfile() error {
	fmt.Println("Create logfile called")
	info, err := os.Lstat(l.LogfilePath)
	if err != nil {
		file, err := os.Create(l.LogfilePath)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	if info.Size() > int64(l.MaxLogfileSize) {
		err := l.ArchiveLogFile()
		if err != nil {
			return err
		}
		file, err := os.Create(l.LogfilePath)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	return nil
}

func (l *Log) ArchiveLogFile() error {
	now := time.Now()
	zipfilename := fmt.Sprintf("Log_%d_%s_%d_%d_%d.zip",
		now.Year(),
		now.Month().String(),
		now.Day(),
		now.Hour(),
		now.Minute())
	archive, err := os.Create(zipfilename)
	if err != nil {
		return fmt.Errorf("error at creating archive-> %s", err.Error())
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()

	logfile, err := os.Open(l.LogfilePath)
	if err != nil {
		return fmt.Errorf("error at opening logfile-> %s", err.Error())
	}
	defer logfile.Close()

	zippedlog, err := zipWriter.Create(l.LogfilePath)
	if err != nil {
		return fmt.Errorf("error at compress logfile-> %s", err.Error())
	}
	if _, err := io.Copy(zippedlog, logfile); err != nil {
		return fmt.Errorf("error at copy compressed users-> %s", err.Error())
	}

	return nil
}

func InitLogger(options ...Option) (*Log, error) {
	logger := Log{
		Logchan:        make(chan string),
		LogfilePath:    "server.log",
		MaxLogfileSize: 1024 * 1024 * 10,
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
