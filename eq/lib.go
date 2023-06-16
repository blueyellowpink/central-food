package eq

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

type ConfluentKafka struct {

}

var _ EventQueue[kafka.Producer, kafka.Consumer] = (*ConfluentKafka)(nil)

// type ConfluentProducerConfig struct {
//     Host string
// }

func (c *ConfluentKafka) NewProducer(cfg map[string]string) (*kafka.Producer, error) {
    producer, err := kafka.NewProducer(&kafka.ConfigMap{
        "bootstrap.servers": cfg["bootstrap.servers"],
    })

    if err != nil {
        return nil, err
    }

    return producer, nil
}

// type ConfluentConsumerConfig struct {
//     Host string
//     GroupId string
// }

func (c *ConfluentKafka) NewConsumer(cfg map[string]string) (*kafka.Consumer, error) {
    consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": cfg["bootstrap.servers"],
        "group.id": cfg["group.id"],
        "auto.offset.reset": cfg["auto.offset.reset"],
    })

    if err != nil {
        return nil, err
    }

    return consumer, nil
}
