package logging

import (
	"log"
	"os"
)

type LoggerDefault struct {
	logger    *log.Logger
	LOG_LEVEL string
}

func InitDefault() LoggerImpl {
	l := LoggerDefault{
		logger:    log.New(os.Stdout, os.Getenv("LOG_LEVEL")+": ", log.Ldate|log.Ltime),
		LOG_LEVEL: os.Getenv("LOG_LEVEL"),
	}

	return &l
}

func (l *LoggerDefault) Info(v ...interface{}) {
	if l.LOG_LEVEL != "INFO" {
		return
	}

	data := toData(v)
	l.logger.Print(data)
}

func (l *LoggerDefault) Error(v ...interface{}) {
	if l.LOG_LEVEL != "ERROR" {
		return
	}

	data := toData(v)
	l.logger.Print(data)
}

func (l *LoggerDefault) Debug(v ...interface{}) {
	if l.LOG_LEVEL != "DEBUG" {
		return
	}

	data := toData(v)
	l.logger.Print(data)
}
