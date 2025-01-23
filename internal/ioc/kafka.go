package ioc

import (
	"github.com/IBM/sarama"
	"github.com/raiki02/EG/internal/cache"
	"github.com/spf13/viper"
)

// comment 和 like需要操作，嵌入act，post中
type KafkaHdl interface {
	Produce(topic string, msg []byte) error
	Consume(topic string) error
}

type Kafka struct {
	consumer sarama.Consumer
	producer sarama.SyncProducer
	cache    cache.CacheHdl
}

func NewKafka(c cache.CacheHdl) *Kafka {
	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Return.Successes = true
	conf.Producer.Return.Errors = true
	producer, _ := sarama.NewSyncProducer([]string{viper.GetString("kafka.host")}, conf)

	consumer, _ := sarama.NewConsumer([]string{viper.GetString("kafka.host")}, nil)
	return &Kafka{producer: producer,
		consumer: consumer,
		cache:    c,
	}
}

// 有赞和评论两topic
func (k *Kafka) Produce(topic string, msg []byte) error {
	_, _, err := k.producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	})
	return err
}

// 消耗之后，将数据存入
func (k *Kafka) Consume(topic string) {
	partitionConsumer, err := k.consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			// 存入数据库
			println(string(msg.Value))
		}
	}
}
