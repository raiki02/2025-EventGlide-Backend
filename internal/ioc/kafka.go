package ioc

import "github.com/IBM/sarama"

type KafkaHdl interface {
	Produce(topic string, key string, value []byte) error
	Consume(topic string, group string, handler func([]byte) error) error
}

type Kafka struct {
	producer *sarama.AsyncProducer
	consumer *sarama.ConsumerGroup
}

func NewKafkaHdl() *Kafka {
	return &Kafka{}
}
