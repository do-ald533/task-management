package logging

import (
	"encoding/json"
	"os"
)

// Logger is logging interface.
type LoggerImpl interface {
	Info(v ...interface{})
	Debug(v ...interface{})
	Error(v ...interface{})
}

func Info(v ...interface{}) {
	Init().Info(v...)
}

func Debug(v ...interface{}) {
	Init().Debug(v...)
}

func Error(v ...interface{}) {
	Init().Error(v...)
}

func Init() LoggerImpl {
	var log LoggerImpl = InitDefault()

	if os.Getenv("CLOUDWATCH_LOG") == "true" {
		log = InitCloudWatch()
	}

	return log
}

func toData(v []interface{}) string {
	LOG_LEVEL := os.Getenv("LOG_LEVEL")

	var rawdata interface{} = map[string]interface{}{
		"level": LOG_LEVEL,
		"trace": longFileName(),
		"data":  v,
	}

	data, _ := json.MarshalIndent(rawdata, "", "  ")

	return string(data)
}
