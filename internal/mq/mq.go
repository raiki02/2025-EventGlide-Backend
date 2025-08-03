package mq

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
)

type MQHdl interface {
}

type MQ struct {
	rdb *redis.Client
}

func NewMQ(rdb *redis.Client) *MQ {
	mq := &MQ{
		rdb: rdb,
	}
	return mq
}

func (mq *MQ) Publish(ctx context.Context, channel string, message interface{}) error {
	jsonReq, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return mq.rdb.Publish(ctx, channel, jsonReq).Err()
}

func (mq *MQ) Subscribe(ctx context.Context, channel string) *redis.PubSub {
	return mq.rdb.Subscribe(ctx, channel)
}
