package engine

import "sync"

type Middleware struct {
	handlers []HandlerFunc
	mu       sync.RWMutex
}

func NewMiddleware() *Middleware {
	return &Middleware{
		handlers: make([]HandlerFunc, 0),
	}
}

func (m *Middleware) Add(handler HandlerFunc) {
	m.mu.Lock()
	defer func() {
		m.mu.Unlock()
	}()
	m.handlers = append(m.handlers, handler)
}
