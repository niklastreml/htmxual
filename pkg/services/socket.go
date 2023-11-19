package services

import (
	"sync"
	"time"
)

func NewSocketService() *SocketService {
	return &SocketService{
		messages:    []Message{},
		mu:          sync.Mutex{},
		subscribers: map[chan Message]bool{},
	}

}

type SocketService struct {
	messages []Message
	mu       sync.Mutex

	subscribers map[chan Message]bool
}

type Message struct {
	Timestamp time.Time
	User      string
	Body      string
}

func (ss *SocketService) Subscribe() chan Message {
	c := make(chan Message)
	ss.subscribers[c] = true
	return c
}

func (ss *SocketService) Unsubscribe(c chan Message) {
	close(c)
	delete(ss.subscribers, c)
}

func (ss *SocketService) Publish(m Message) {
	m.Timestamp = time.Now()
	ss.mu.Lock()
	defer ss.mu.Unlock()
	ss.messages = append(ss.messages, m)

	for sub := range ss.subscribers {
		sub <- m
	}
}

func (ss *SocketService) Messages() []Message {
	return ss.messages
}
