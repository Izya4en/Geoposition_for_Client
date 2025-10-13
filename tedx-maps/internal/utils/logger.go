package utils

import (
	"log"
	"os"
)

type Logger struct {
	info  *log.Logger
	error *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		info:  log.New(os.Stdout, "[INFO] ", log.LstdFlags|log.Lshortfile),
		error: log.New(os.Stderr, "[ERROR] ", log.LstdFlags|log.Lshortfile),
	}
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.error.Println(v...)
}
