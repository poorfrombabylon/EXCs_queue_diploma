package queue_converter

import (
	"context"
	"excs_queue/internal/config"
	"fmt"
	"github.com/go-redis/redis"
)

type QueueConverter interface {
	SubscribeToSubscriberChannel(_ context.Context) *redis.PubSub
}

type queueConverter struct {
	client         *redis.Client
	topic          string
	deliveryMethod string
}

func NewQueueConverter(cfg config.QueueRedisConfig) (QueueConverter, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", cfg.Host, cfg.Port),
		Password: cfg.Password,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &queueConverter{
		client:         client,
		topic:          cfg.Topic,
		deliveryMethod: cfg.DeliveryMethod,
	}, nil
}

func (c *queueConverter) SubscribeToSubscriberChannel(_ context.Context) *redis.PubSub {
	return c.client.Subscribe(c.topic)
}
