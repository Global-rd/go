package logger

import (
	"log/slog"
	"os"
)

type Option interface {
	apply(*Log)
}

type OptionFunc func(*Log)

func (f OptionFunc) apply(logger *Log) {
	f(logger)
}

type Log struct {
	Loghandler     *slog.Logger
	LogfilePath    string
	MaxLogfileSize int
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

func (l Log) INFO(info string) error {
	logFile, err := os.OpenFile(l.LogfilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		slog.Error("failed to open log file", "error", err)
		return err
	}
	defer logFile.Close()
	l.Loghandler.Info(info)
	return nil
}

func (l Log) WARNING(info string) error {
	logFile, err := os.OpenFile(l.LogfilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		slog.Warn("failed to open log file", "error", err)
		return err
	}
	defer logFile.Close()
	l.Loghandler.Info(info)
	return nil
}

func (l Log) ERROR(info string) error {
	logFile, err := os.OpenFile(l.LogfilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		slog.Error("failed to open log file", "error", err)
		return err
	}
	defer logFile.Close()
	l.Loghandler.Error(info)
	return nil
}

func InitLogger(options ...Option) (*Log, error) {

	logger := Log{
		LogfilePath:    "server.log",
		MaxLogfileSize: 1024 * 1024 * 10,
	}

	for _, option := range options {
		option.apply(&logger)
	}

	logFile, err := os.OpenFile(logger.LogfilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		slog.Error("failed to open log file", "error", err)
		return &logger, err
	}
	defer logFile.Close()

	// Create a logger that writes to the file
	l := slog.New(slog.NewTextHandler(logFile, nil))
	logger.Loghandler = l

	l.Info("Logger started")

	return &logger, nil

}
