package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	//log.SetLevel(log.WarnLevel)

	log.SetOutput(os.Stdout)

	log.WithField("name","aaa").Info(2222)
	log.WithFields(log.Fields{"name":"bbb"}).Info(2222)
}
