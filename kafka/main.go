package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的信息将在success channel返回

	msg := &sarama.ProducerMessage{
		Topic: "topic",
		Value: sarama.StringEncoder("first send msg"),
	}
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:2222"}, config)
	if err != nil {
		return
	}
	defer client.Close()

	pid, offset, err := client.SendMessage(msg)
	fmt.Println(pid, offset, err)
}
