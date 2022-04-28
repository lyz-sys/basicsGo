package logger

import (
	"demo/config"
	"log"
	"os"
)

var Logger *log.Logger

func init() {
	logFile, err := os.OpenFile(config.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	Logger = log.New(logFile, "", log.Llongfile)
}
