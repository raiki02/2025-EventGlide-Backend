package ioc

import (
	"github.com/IBM/sarama"
	"github.com/spf13/viper"
)

type KafkaHdl interface {
	Produce(topic string, key string, value []byte) error
	Consume(topic string, group string, handler func([]byte) error) error
}

type Kafka struct {
	producer *sarama.AsyncProducer
	consumer *sarama.ConsumerGroup
}

func NewKafka(producer *sarama.AsyncProducer, consumer *sarama.ConsumerGroup) *Kafka {
	return &Kafka{
		producer: producer,
		consumer: consumer,
	}
}

func NewAsyncProducer() (*sarama.AsyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	addr := viper.GetString("kafka.addr")
	producer, err := sarama.NewAsyncProducer([]string{addr}, config)
	if err != nil {
		return nil, err
	}
	return &producer, nil
}

func NewConsumerGroup() (*sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()
	addr := viper.GetString("kafka.addr")
	consumer, err := sarama.NewConsumerGroup([]string{addr}, "EG-group", config)
	if err != nil {
		return nil, err
	}
	return &consumer, nil
}
