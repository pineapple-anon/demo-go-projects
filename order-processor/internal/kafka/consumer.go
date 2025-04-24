package kafka

import (
	"github.com/segmentio/kafka-go"
)

// Consumer represents a Kafka consumer
func NewConsumer(topic, broker string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: "order-consumer-group",
	})
} 