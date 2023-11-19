package services

import (
	"sync"
	"time"
)

func NewSseService() *SseService {
	return &SseService{
		messages:    []Message{},
		mu:          sync.Mutex{},
		subscribers: map[chan Message]bool{},
	}

}

type SseService struct {
	messages []Message
	mu       sync.Mutex

	subscribers map[chan Message]bool
}

type Message struct {
	Timestamp time.Time
	User      string
	Body      string
}

func (ss *SseService) Subscribe() chan Message {
	c := make(chan Message)
	ss.subscribers[c] = true
	return c
}

func (ss *SseService) Unsubscribe(c chan Message) {
	close(c)
	delete(ss.subscribers, c)
}

func (ss *SseService) Publish(m Message) {
	m.Timestamp = time.Now()
	ss.mu.Lock()
	defer ss.mu.Unlock()
	ss.messages = append(ss.messages, m)

	for sub := range ss.subscribers {
		sub <- m
	}
}

func (ss *SseService) Messages() []Message {
	return ss.messages
}
