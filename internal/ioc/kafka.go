package ioc

import (
	"github.com/IBM/sarama"
)

type Kafka struct {
	P sarama.SyncProducer
	C sarama.Consumer
}

func NewClient(addr []string) sarama.Client {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	client, err := sarama.NewClient(addr, config)
	if err != nil {
		panic(err)
	}
	return client
}

func NewProducer(client sarama.Client) sarama.SyncProducer {
	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		panic(err)
	}
	return producer
}

func NewConsumer(client sarama.Client) sarama.Consumer {
	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		panic(err)
	}
	return consumer
}

func NewKafka(p sarama.SyncProducer, c sarama.Consumer) *Kafka {
	return &Kafka{
		P: p,
		C: c,
	}
}
