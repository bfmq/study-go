package main

import (
	"fmt"
	"os"

	// "sync"

	"gitlab.renrenchongdian.com/studygo/logtransfer/conf"
	"gitlab.renrenchongdian.com/studygo/logtransfer/es"
	"gitlab.renrenchongdian.com/studygo/logtransfer/kafka"
	"gopkg.in/ini.v1"
)

var (
	cfg         = new(conf.LogTransferCfg)
	basePath, _ = os.Getwd()
	// wg          sync.WaitGroup
)

func main() {
	// 加载配置文件
	err := ini.MapTo(cfg, basePath+`\conf\cfg.ini`)
	if err != nil {
		fmt.Printf("load config failed,err:%s\n", err)
		return
	}

	// 初始化kafka
	err = kafka.Init(cfg.KafkaCfg.Addr, cfg.KafkaCfg.Topic)
	if err != nil {
		fmt.Printf("kafka init failed,err:%s\n", err)
		return
	}

	// 初始化es
	err = es.Init(cfg.ESCfg.Addr, cfg.ESCfg.ChanSize, cfg.ESCfg.Nums)
	if err != nil {
		fmt.Printf("es init failed,err:%s\n", err)
		return
	}

	// 发给es

}
