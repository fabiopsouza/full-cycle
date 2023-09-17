package main

import (
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

func main() {
	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "fcutils",
	}

	c, err := ckafka.NewConsumer(&configMap)
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening for balance events...")
	c.SubscribeTopics([]string{"balance"}, nil)
	run := true
	for run {
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else if err.(ckafka.Error).IsFatal() {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			run = false
		}
	}
	c.Close()
}
