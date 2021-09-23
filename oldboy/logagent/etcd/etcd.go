package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	client *clientv3.Client
)

type LogEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

// 初始化client
func Init(addrs []string, timeout time.Duration) (err error) {
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   addrs,
		DialTimeout: timeout,
	})
	if err != nil {
		// handle error!
		fmt.Printf("init etcd failed, err:%v\n", err)
		return
	}

	return
}

// 从etcd中获取值
func GetConf(key string) (logEntryConf []*LogEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := client.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		err = json.Unmarshal(ev.Value, &logEntryConf)
		if err != nil {
			fmt.Printf("unmarshal etcd value failed,err:%s\n", err)
			return
		}
	}
	return
}

func WatchConf(key string, newConfChan chan []*LogEntry) {
	ch := client.Watch(context.Background(), key)
	for wrep := range ch {
		for _, evt := range wrep.Events {
			var newConf []*LogEntry
			err := json.Unmarshal(evt.Kv.Value, &newConf)
			if err != nil {
				fmt.Printf("unmarshal failed,err:%s\n", err)
				continue
			}
			fmt.Println("get new conf !", newConf)
			newConfChan <- newConf
		}
	}
}
