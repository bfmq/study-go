package main

import (
	"gitlab.renrenchongdian.com/studygo/day08/mylogger"
)

var log mylogger.Logger

func main() {
	// log = mylogger.NewConsoleLog("error")
	log = mylogger.NewFileLog("info", `C:\Users\Aone\go\src\gitlab.renrenchongdian.com\studygo\day08\05mylogger_test`, "run.log", 10*1024*1024)

	for {
		log.Debug("这是一条debug日志")
		log.Info("这是一条info日志")
		log.Warning("这是一条warning日志")
		log.Error("这是一条error日志")
		log.Fatal("这是一条fatal日志")
		// time.Sleep(time.Second)
	}
}
