package kafka

import "github.com/segmentio/kafka-go"

type Tool interface {
	CreateReader(topic string, GroupID string) *kafka.Reader
	CreateWriter(topic string) *kafka.Writer
	CreateConn(topic string) *kafka.Conn
}
