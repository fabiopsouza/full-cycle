package kafka

import (
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

type Consumer struct {
	ConfigMap *ckafka.ConfigMap
	Topics    []string
}

func NewConsumer(configMap *ckafka.ConfigMap, topics []string) *Consumer {
	return &Consumer{
		ConfigMap: configMap,
		Topics:    topics,
	}
}

func (c *Consumer) Consume(Handle func(value []byte) error) {
	consumer, err := ckafka.NewConsumer(c.ConfigMap)
	if err != nil {
		panic(err)
	}

	err = consumer.SubscribeTopics(c.Topics, nil)
	if err != nil {
		panic(err)
	}

	run := true
	fmt.Println("Listening for balance events...")
	for run {
		msg, err := consumer.ReadMessage(time.Second)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			err = Handle(msg.Value)
			if err != nil {
				fmt.Printf("Error on handle event: %v\n", err)
			}
		} else if err.(ckafka.Error).IsFatal() {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			run = false
		}
	}
	err = consumer.Close()
	if err != nil {
		panic(err)
	}
}
