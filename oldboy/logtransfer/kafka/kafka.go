package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
	"gitlab.renrenchongdian.com/studygo/logtransfer/es"
)

// 初始化kafka消费者
func Init(addrs string, topic string) (err error) {
	consumer, err := sarama.NewConsumer([]string{addrs}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}

	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return err
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
				ld := &es.LogData{
					Data:  string(msg.Value),
					Topic: topic,
				}
				es.SentToESChan(ld)
			}
		}(pc)
	}
	return
}
