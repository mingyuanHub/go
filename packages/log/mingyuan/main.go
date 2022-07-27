package main

import (
	"os"
	"fmt"
	"sync"
	"time"
)

var (
	Hour        = -1
	SdkEvV2Path = fmt.Sprintf("./%s/sdk-ev.log", LogsDirV2)
	Lock        sync.RWMutex
	FileWriter  *os.File
	LogsDirV2   = "logs_v2" //日志目录
)

func getFileWriter() (fileWriter *os.File, err error) {

	timeLocation := time.FixedZone("UTC", 8*3600)
	nowTime := time.Now().Add(-time.Hour).In(timeLocation)
	day := nowTime.Format("2006-01-02")
	hour := nowTime.Hour()

	if Hour == -1 {
		Hour = hour
		err = openEvV2File()
		return FileWriter, err
	}

	Lock.RLock()
	if hour != Hour {
		Lock.RUnlock()
		Lock.Lock()
		if hour != Hour {
			FileWriter.Close()
			newFileName := fmt.Sprintf("./%s/%s-%d-sdk-ev.log", LogsDirV2, day, hour)
			err := os.Rename(SdkEvV2Path, newFileName)
			if err == nil {
				Hour = hour
				openEvV2File()
			}
		}
		Lock.Unlock()
	} else {
		Lock.RUnlock()
	}

	return FileWriter, nil
}

func openEvV2File() (error) {
	fileWriter, err := os.OpenFile(SdkEvV2Path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	FileWriter = fileWriter
	return nil
}