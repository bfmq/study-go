package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"gopkg.in/ini.v1"

	"gitlab.renrenchongdian.com/studygo/logagent/conf"
	"gitlab.renrenchongdian.com/studygo/logagent/etcd"
	"gitlab.renrenchongdian.com/studygo/logagent/kafka"
	"gitlab.renrenchongdian.com/studygo/logagent/taillog"
	"gitlab.renrenchongdian.com/studygo/logagent/utils"
)

// agent入口程序

var (
	cfg         = new(conf.AppConf)
	basePath, _ = os.Getwd()
	wg          sync.WaitGroup
)

func main() {
	// 加载配置文件
	// err := ini.MapTo(cfg, `c:\Users\Aone\go\src\gitlab.renrenchongdian.com\studygo\logagent\conf\config.ini`)
	err := ini.MapTo(cfg, basePath+`\conf\config.ini`)
	if err != nil {
		fmt.Printf("load config failed,err:%s\n", err)
		return
	}

	// kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Addr}, cfg.KafkaConf.ChanMaxSize)
	if err != nil {
		fmt.Printf("init kafka failed,err:%s\n", err)
		return
	}
	fmt.Println("init kafka success.")

	// 初始化etcd
	err = etcd.Init([]string{cfg.EtcdConf.Addr}, time.Duration(cfg.EtcdConf.TimeOut)*time.Second)
	if err != nil {
		fmt.Printf("init etcd failed,err:%s\n", err)
		return
	}
	fmt.Println("init etcd success.")

	// 获取自己的专用key
	ipStr, err := utils.GetOutBoundIP()
	if err != nil {
		fmt.Printf("get ip failed,err:%s\n", err)
		return
	}
	etcdConfKey := fmt.Sprintf(cfg.EtcdConf.Key, ipStr)

	// 获取要收集日志配置
	logEntryConf, err := etcd.GetConf(etcdConfKey)
	if err != nil {
		fmt.Printf("etcd.GetConf failed,err:%s\n", err)
		return
	}
	fmt.Printf("get conf from etcd success, %v\n", logEntryConf)

	// 开日志收集
	taillog.Init(logEntryConf)

	// 开配置文件监听
	newConfChan := taillog.NewConfChan()
	wg.Add(1)
	go etcd.WatchConf(cfg.EtcdConf.Key, newConfChan)
	wg.Wait()
}
