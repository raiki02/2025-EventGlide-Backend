package main

import "github.com/IBM/sarama"

type Kafka struct {
	P *sarama.AsyncProducer
	C *sarama.Consumer
}

func NewProducer() *sarama.AsyncProducer {
	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true
	conf.Producer.Return.Errors = true
	conf.Producer.RequiredAcks = sarama.WaitForAll
	p, _ := sarama.NewAsyncProducer([]string{"localhost:9092"}, conf)
	return &p
}

func NewConsumer() *sarama.Consumer {
	conf := sarama.NewConfig()
	conf.Consumer.Return.Errors = true
	c, _ := sarama.NewConsumer([]string{"localhost:9092"}, conf)
	return &c
}

var K *Kafka

func NewKafka() *Kafka {
	return &Kafka{
		P: NewProducer(),
		C: NewConsumer(),
	}
}

type Number struct {
	Sid string
	Bid string
}

// produce
func (k *Kafka) AddLikesNum() {
	n := Number{
		Sid: "123",
		Bid: "456",
	}
	(*k.P).Input() <- &sarama.ProducerMessage{
		Topic: "likes",
		Value: sarama.StringEncoder(n.Sid + ":" + n.Bid),
	}
}

func CutLikesNum() {
	n := Number{
		Sid: "123",
		Bid: "456",
	}
	(*K.P).Input() <- &sarama.ProducerMessage{
		Topic: "likes",
		Value: sarama.StringEncoder(n.Sid + ":" + n.Bid),
	}
}

func AddCommentsNum() {
	n := Number{
		Sid: "123",
		Bid: "456",
	}
	(*K.P).Input() <- &sarama.ProducerMessage{
		Topic: "comments",
		Value: sarama.StringEncoder(n.Sid + ":" + n.Bid),
	}
}

// consume
