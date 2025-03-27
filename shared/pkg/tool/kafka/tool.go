package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

type tool struct {
}

func New() Tool {
	return &tool{}
}

func (t *tool) CreateReader(topic string, GroupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		Topic:       topic,
		Partition:   0,
		MaxBytes:    10e6, // 10MB
		MaxAttempts: 5,
		StartOffset: kafka.FirstOffset,
		GroupID:     GroupID,
	})
}

func (t *tool) CreateWriter(topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:  kafka.TCP([]string{"localhost:9092"}...),
		Topic: topic,
		Async: true,
	}
}

func (t *tool) CreateConn(topic string) *kafka.Conn {
	// 創建 conn
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, 0)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return conn
}
