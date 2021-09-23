package es

import (
	"context"
	"strings"
	"time"

	"github.com/olivere/elastic/v7"
)

type LogData struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

var (
	client *elastic.Client
	ch     chan *LogData
)

func Init(addr string, chanSize int, nums int) (err error) {
	if !strings.HasPrefix(addr, "http://") {
		addr = "http://" + addr
	}

	client, err = elastic.NewClient(elastic.SetURL(addr))
	if err != nil {
		return
	}
	ch = make(chan *LogData, chanSize)
	for i := 0; i < nums; i++ {
		go SendToES()
	}

	return
}

func SentToESChan(msg *LogData) {
	ch <- msg
}

func SendToES() {
	for {
		select {
		case msg := <-ch:
			_, err := client.Index().
				Index(msg.Topic).
				BodyJson(msg).
				Do(context.Background())
			if err != nil {
				continue
			}
		default:
			time.Sleep(time.Millisecond * 500)
		}
	}
}
