package engine

import (
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

type Client struct {
	conn   *websocket.Conn // 當前連線
	engine *Engine
	keys   map[string]interface{}
	mu     sync.RWMutex
	send   chan []byte
}

// 監聽用戶寫入指令
func (c *Client) run() {
	go c.read()
	go c.write()
}

func (c *Client) Conn() *websocket.Conn {
	return c.conn
}

func (c *Client) read() {
	defer func() {
		c.engine.channelManager.DelAll(c)
		_ = c.conn.Close()
	}()
	for {
		// 讀取客戶端消息(阻塞等待)
		_, b, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("error:", err)
			break
		}
		// 調用路由解析器
		path := c.engine.resolver(b)
		// 路由選擇
		handlers, ok := c.engine.route.Get(path)
		if !ok {
			continue
		}
		// 創建 context
		context := Context{
			engine:   c.engine,
			client:   c,
			handlers: handlers,
			keys:     make(map[string]interface{}),
			index:    -1,
			data:     b,
		}
		// 執行路由
		go context.Next()
	}
}

// 監聽Hub發來的指令
func (c *Client) write() {
	defer func() {
		c.engine.channelManager.DelAll(c)
		_ = c.conn.Close()
	}()
	for {
		select {
		case b, ok := <-c.send:
			if !ok {
				break
			}
			_ = c.conn.WriteMessage(websocket.BinaryMessage, b)
		}
	}
}

func (c *Client) Set(key string, value interface{}) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	c.keys[key] = value
}

func (c *Client) Get(key string) (value interface{}, exists bool) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	value, exists = c.keys[key]
	return
}

func (c *Client) Del(key string) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	delete(c.keys, key)
}

func (c *Client) MustGet(key string) interface{} {
	if value, exists := c.Get(key); exists {
		return value
	}
	panic("Key \"" + key + "\" does not exist")
}

func (c *Client) Send(msg []byte) {
	c.send <- msg
}
