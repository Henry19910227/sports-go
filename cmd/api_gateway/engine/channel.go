package engine

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/segmentio/kafka-go"
	"sync"
)

type channelManager struct {
	channels map[string]*channel
	mu       sync.RWMutex
}

func (c *channelManager) createGroup(name string) {
	_, ok := c.channels[name]
	if ok {
		return
	}
	//ch := NewChannel(name, kafkaReader(name, c.kafkaSetting), kafkaWriter(name, c.kafkaSetting))
	//go ch.Run()
	ch := NewChannel(name, nil, nil)
	c.channels[name] = ch
}

func (c *channelManager) Add(name string, client *Client) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	_, ok := c.channels[name]
	if !ok {
		c.createGroup(name)
	}
	c.channels[name].AddClient(client)
}

func (c *channelManager) Del(name string, client *Client) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	c.channels[name].DelClient(client)
}

func (c *channelManager) DelAll(client *Client) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	for _, ch := range c.channels {
		ch.DelClient(client)
	}
}

func (c *channelManager) Clients(channel string) map[*Client]*Client {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	if _, ok := c.channels[channel]; !ok {
		fmt.Println("channel not found")
		return nil
	}
	return c.channels[channel].clients
}

func (c *channelManager) Send(channel string, b []byte) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	if _, ok := c.channels[channel]; !ok {
		return
	}
	for _, client := range c.channels[channel].Clients() {
		_ = client.Conn().WriteMessage(websocket.BinaryMessage, b)
	}
}

type channel struct {
	reader  *kafka.Reader
	writer  *kafka.Writer
	clients map[*Client]*Client
	name    string
	mu      sync.RWMutex
}

func NewChannel(name string, reader *kafka.Reader, writer *kafka.Writer) *channel {
	return &channel{
		reader:  reader,
		writer:  writer,
		name:    name,
		clients: make(map[*Client]*Client),
	}
}

func (c *channel) AddClient(client *Client) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	c.clients[client] = client
}

func (c *channel) DelClient(client *Client) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	delete(c.clients, client)
}

func (c *channel) Clients() map[*Client]*Client {
	return c.clients
}

func (c *channel) Send(data []byte) {
	_ = c.writer.WriteMessages(context.Background(),
		kafka.Message{Value: data},
	)
}

func (c *channel) Run() {
	defer func() {
		_ = c.reader.Close()
	}()
	for {
		m, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			break
		}
		for _, client := range c.clients {
			client.send <- m.Value
		}
	}
}
