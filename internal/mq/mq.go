package mq

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type MQHdl interface {
	Publish(ctx context.Context, stream string, message interface{}) error
	Consume(ctx context.Context, stream, lastIDKey string, count int64, block time.Duration) ([]redis.XMessage, error)
}

type MQ struct {
	rdb *redis.Client
}

func NewMQ(rdb *redis.Client) MQHdl {
	mq := &MQ{
		rdb: rdb,
	}
	return mq
}

func (mq *MQ) Publish(ctx context.Context, stream string, message interface{}) error {
	jsonReq, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return mq.rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: stream,
		Values: map[string]interface{}{
			"data": jsonReq,
		},
	}).Err()
}

func (mq *MQ) Consume(ctx context.Context, stream, lastIDKey string, count int64, block time.Duration) ([]redis.XMessage, error) {
	lastID, err := mq.rdb.Get(ctx, lastIDKey).Result()
	if err == redis.Nil {
		lastID = "0"
	} else if err != nil {
		return nil, err
	}

	res, err := mq.rdb.XRead(ctx, &redis.XReadArgs{
		Streams: []string{stream, lastID},
		Count:   count,
		Block:   block,
	}).Result()

	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	if len(res) == 0 || len(res[0].Messages) == 0 {
		return nil, nil
	}

	lastID = res[0].Messages[len(res[0].Messages)-1].ID
	_ = mq.rdb.Set(ctx, lastIDKey, lastID, 0).Err()

	return res[0].Messages, nil
}
