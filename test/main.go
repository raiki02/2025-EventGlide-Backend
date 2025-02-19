package main

import (
	"github.com/IBM/sarama"
	"log"
	"time"
)

func main() {
	conf := sarama.NewConfig()

	conf.Producer.Return.Successes = true
	conf.Producer.Partitioner = sarama.NewRandomPartitioner
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Consumer.Return.Errors = true

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, conf)
	log.Println("Producer created")
	if err != nil {
		panic(err)
	}
	defer producer.Close()
	produceMessage(producer)
	produceMessage(producer)
	produceMessage(producer)
	produceMessage(producer)
	produceMessage(producer)
	log.Println("Messages sent")

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	log.Println("Consumer created")
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("test", 0, sarama.OffsetNewest)
	log.Println("Partition consumer created")
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Println("Received messages", string(msg.Key), string(msg.Value))
		case err := <-partitionConsumer.Errors():
			log.Println("Received error", err.Error())
		default:
			time.Sleep(1 * time.Second)
			log.Println("No message received, waiting for message")
		}
	}

}

func produceMessage(producer sarama.SyncProducer) {

	msg := &sarama.ProducerMessage{
		Topic: "test",
		Key:   sarama.StringEncoder("key"),
		Value: sarama.StringEncoder("value"),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Println("Failed to send message", err)
	} else {
		log.Println("Message sent to partition", partition, "offset", offset)
	}
}
