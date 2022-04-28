package config

import (
	"log"
	"os"
	"strings"

	"github.com/go-ini/ini"
)

var (
	a               *ini.File
	Port            int
	AppKey          string
	KafkaBrokerList []string
	err             error
	LogFile         string
)

func init() {
	a, err = ini.Load(os.Getenv("GOPATH") + "/src/demo/config/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'config/app.ini': %v", err)
	}
	appInit()
	kafkaInit()
}

func appInit() {
	app := a.Section("APP")
	AppKey = app.Key("APP_KEY").MustString("")
	Port = app.Key("APP_PORT").MustInt(80)
	LogFile = app.Key("LOG_FILE").MustString("")
}

func kafkaInit() {
	conf := a.Section("KAFKA")
	kafka := conf.Key("BROKER_LIST").MustString("")

	KafkaBrokerList = strings.Split(kafka, ",")
}
