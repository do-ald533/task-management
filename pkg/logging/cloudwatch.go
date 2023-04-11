package logging

import (
	"fmt"
	"goapi/pkg/aws/cloudwatch"
	"log"
	"os"
	"runtime"
)

type LoggerCloudWatch struct {
	logger    *log.Logger
	LOG_LEVEL string
}

func InitCloudWatch() LoggerImpl {
	l := LoggerCloudWatch{
		logger:    log.New(os.Stdout, os.Getenv("LOG_LEVEL")+": ", log.Ldate|log.Ltime),
		LOG_LEVEL: os.Getenv("LOG_LEVEL"),
	}

	return &l
}

func longFileName() string {
	_, file, line, ok := runtime.Caller(4)

	if ok {
		return fmt.Sprintf("%s:%d", file, line)
	}

	return ""
}

func (l *LoggerCloudWatch) Info(v ...interface{}) {
	if l.LOG_LEVEL != "INFO" {
		return
	}

	data := toData(v)

	service := cloudwatch.Init()
	go service.Log(&data)

	l.logger.Print(data)
}

func (l *LoggerCloudWatch) Error(v ...interface{}) {
	if l.LOG_LEVEL != "ERROR" {
		return
	}

	data := toData(v)

	service := cloudwatch.Init()
	go service.Log(&data)

	l.logger.Print(data)
}

func (l *LoggerCloudWatch) Debug(v ...interface{}) {
	if l.LOG_LEVEL != "DEBUG" {
		return
	}

	data := toData(v)

	service := cloudwatch.Init()
	go service.Log(&data)

	l.logger.Print(data)
}
