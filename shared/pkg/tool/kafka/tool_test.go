package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"testing"
	"time"
)

func TestCreateTopic(t *testing.T) {
	// 創建 conn
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "conn1", 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	conn1, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "conn2", 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	// 寫入資料
	_ = conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	_ = conn1.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn1.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	// 關閉連接
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

	if err := conn1.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func TestConsume(t *testing.T) {
	// to consume messages
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "my-topic", 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	_ = conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))
	}

	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}

func TestWriter(t *testing.T) {
	w := &kafka.Writer{
		Addr:  kafka.TCP([]string{"localhost:9092"}...),
		Topic: "Hello",
	}
	err := w.WriteMessages(context.Background(),
		kafka.Message{Value: []byte("one!")},
	)
	log.Fatal(err)
}

func TestReader(t *testing.T) {
	// 設置超時時間
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// make a new reader that consumes from topic-A, partition 0, at offset 42
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		Topic:       "Hello",
		Partition:   0,
		MaxBytes:    10e6, // 10MB
		StartOffset: kafka.FirstOffset,
		GroupID:     "1009",
	})
	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))

		// 確認消息已讀取（可選）
		if err := r.CommitMessages(ctx, m); err != nil {
			log.Fatalf("Failed to commit message: %v", err)
		}
	}
	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

func TestListTopic(t *testing.T) {
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	for k := range m {
		fmt.Println(k)
	}
}
