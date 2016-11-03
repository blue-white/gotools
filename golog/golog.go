package golog

import (
	"fmt"
	"log"
	"os"
)

const (
	// log level
	OFF   = -1
	Fatal = iota
	Error
	Info
	Debug
)

type Logger struct {
	fatal, err, info, debug *log.Logger
}

func NewLogger(level int8) *Logger {
	logger := &Logger{}

	switch {
	case level >= Fatal:
		logger.fatal = log.New(os.Stderr, "\033[0;35mFATAL:\033[0m", log.Ldate|log.Ltime|log.Lshortfile)
		fallthrough
	case level >= Error:
		logger.err = log.New(os.Stderr, "\033[0;31mERROR:\033[0m", log.Ldate|log.Ltime|log.Lshortfile)
		fallthrough
	case level >= Info:
		logger.info = log.New(os.Stdout, "\033[0;33mINFO:\033[0m", log.Ldate|log.Ltime|log.Lshortfile)
		fallthrough
	case level >= Debug:
		logger.debug = log.New(os.Stdout, "\033[0;36mDEBUG:\033[0m", log.Ldate|log.Ltime|log.Lshortfile)
	}
	return logger
}

// Fatal outputs standard error message and terminates program
func (l *Logger) Fatal(format string, v ...interface{}) {
	if l.fatal != nil {
		l.fatal.Output(2, fmt.Sprintf(format, v...))
		os.Exit(1)
	}
}

// Error outputs standard error message
func (l *Logger) Error(format string, v ...interface{}) {
	if l.err != nil {
		l.err.Output(2, fmt.Sprintf(format, v...))
	}
}

// Info outputs standard out message
func (l *Logger) Info(format string, v ...interface{}) {
	if l.info != nil {
		l.info.Output(2, fmt.Sprintf(format, v...))
	}
}

// Debug outputs standard out message
func (l *Logger) Debug(format string, v ...interface{}) {
	if l.debug != nil {
		l.debug.Output(2, fmt.Sprintf(format, v...))
	}
}
