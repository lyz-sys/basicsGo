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
	KafkaBrokerList []string
	err             error
)

func init() {
	a, err = ini.Load(os.Getenv("GOPATH") + "/src/demo/config/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'config/app.ini': %v", err)
	}
	kafka := a.Section("KAFKA").Key("BROKER_LIST").MustString("")
	Port, err = a.Section("APP").Key("APP_PORT").Int()
	if err != nil {
		log.Fatalf("Fail to parse 'config/app.ini': %v", err)
	}
	KafkaBrokerList = strings.Split(kafka, ",")
}
