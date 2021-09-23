package main

import (
	"context"
	"fmt"
	"time"

 	"go.etcd.io/etcd/client/v3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"10.10.2.20:3379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	value := `[{"path":"C:\\Users\\Aone\\go\\src\\gitlab.renrenchongdian.com\\studygo\\logagent\\my.log","topic":"web_log"},{"path":"C:\\Users\\Aone\\go\\src\\gitlab.renrenchongdian.com\\studygo\\logagent\\redis.log","topic":"redis_log"}]`
	// value := `[{"path":"C:\\Users\\Aone\\go\\src\\gitlab.renrenchongdian.com\\studygo\\logagent\\my.log","topic":"web_log"}]`

	_, err = cli.Put(ctx, "/logagent/collect_conf", value)
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
}
