package kafka

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

// 专门向kafka写日志模块

type logData struct {
	topic string
	data  string
}

var (
	client      sarama.SyncProducer
	logDataChan chan *logData
)

// 初始化client
func Init(addrs []string, maxSize int) (err error) {
	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.WaitForAll          //赋值为-1：这意味着producer在follower副本确认接收到数据后才算一次发送完成。
	config.Producer.Partitioner = sarama.NewRandomPartitioner //写到随机分区中，默认设置8个分区
	config.Producer.Return.Successes = true

	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("producer close err, ", err)
		return
	}

	logDataChan = make(chan *logData, maxSize)
	go sendToKafka()
	return
}

func sendToKafka() {
	for {
		select {
		case ld := <-logDataChan:
			msg := &sarama.ProducerMessage{}
			msg.Topic = ld.topic
			msg.Value = sarama.StringEncoder(ld.data)
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				fmt.Println("send message failed, ", err)
				return
			}
			fmt.Printf("分区ID:%v, offset:%v \n", pid, offset)
		default:
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func SendToChan(topic, data string) {
	msg := &logData{
		topic: topic,
		data:  data,
	}
	logDataChan <- msg
}
