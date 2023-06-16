package eq

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

/* Wrapper for Confluent's Kafka Producer */
type ConfluentKafkaProducer struct {
	inner *kafka.Producer
}

var _ Producer[kafka.Producer] = (*ConfluentKafkaProducer)(nil)

func NewConfluentProducer(cfg ConfigMap) (*ConfluentKafkaProducer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": cfg["bootstrap.servers"],
		"acks":              "all",
	})

	if err != nil {
		return nil, err
	}

	return &ConfluentKafkaProducer{inner: producer}, nil
}

func (p *ConfluentKafkaProducer) Inner() *kafka.Producer {
	return p.inner
}

func (p *ConfluentKafkaProducer) Produce(m *Message) error {
	if err := p.inner.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &m.Topic, Partition: kafka.PartitionAny},
		Value:          m.Value,
	}, nil); err != nil {
		return err
	}
	return nil
}

func (p *ConfluentKafkaProducer) Close() {
	p.inner.Close()
}

/* Wrapper for Confluent's Kafka Consumer */
type ConfluentKafkaConsumer struct {
	inner *kafka.Consumer
}

var _ Consumer[kafka.Consumer] = (*ConfluentKafkaConsumer)(nil)

func NewConfluentConsumer(cfg ConfigMap) (*ConfluentKafkaConsumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg["bootstrap.servers"],
		"group.id":          cfg["group.id"],
		"auto.offset.reset": cfg["auto.offset.reset"],
	})

	if err != nil {
		return nil, err
	}

	return &ConfluentKafkaConsumer{inner: consumer}, nil
}

func (c *ConfluentKafkaConsumer) Inner() *kafka.Consumer {
	return c.inner
}

func (c *ConfluentKafkaConsumer) Subscribe(topics []string, handler chan bool) {
	c.inner.SubscribeTopics(topics, nil)

	/* A signal handler or similar could be used to set this to false to break the loop. */
	run := true

	go func() {
		run = <-handler
	}()

	for run {
		msg, err := c.inner.ReadMessage(time.Second)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else if !err.(kafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

func (c *ConfluentKafkaConsumer) Close() {
	c.inner.Close()
}
