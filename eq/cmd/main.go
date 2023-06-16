package main

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	"eq"
)

func producer() {
	/* Confluent's Kafka Producer */
	producer, err := eq.NewConfluentProducer(eq.ConfigMap{
		"bootstrap.servers": "localhost:9094",
	})
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	/* Delivery report handler for produced messages */
	go func() {
		for e := range producer.Inner().Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	/* Produce messages to topic (asynchronously) */
	topic := "test"
	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		if err := producer.Produce(&eq.Message{Topic: topic, Value: []byte(word)}); err != nil {
			panic(err)
		}
	}

	producer.Inner().Flush(15 * 1000)
}

func consumer() {
	/* Confluent's Kafka Consumer */
	consumer, err := eq.NewConfluentConsumer(eq.ConfigMap{
		"bootstrap.servers": "localhost:9094",
		"group.id":          "testGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	run := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		run <- false
	}()
	consumer.Subscribe([]string{"test"}, run)
}

func main() {
	producer()
	consumer()
}
