package main

import (
	"fmt"
	"sync"
	"time"
)
// Message represents a message to be published.
type Message struct {
	Topic   string
	Content interface{}
}

// Subscriber represents a subscriber to a topic.
type Subscriber struct {
	Channel     chan Message
	Unsubscribe chan bool
}

// Broker manages subscribers and publishes messages.
type Broker struct {
	subscribers map[string][]*Subscriber
	mu          sync.RWMutex
}

func NewBroker() *Broker {
	return &Broker{
		subscribers: make(map[string][]*Subscriber),
	}
}

func (b *Broker) Subscribe(topic string) *Subscriber {
    b.mu.Lock()
	subscriber := &Subscriber{
		Channel:     make(chan Message),
		Unsubscribe: make(chan bool),
	}

	
	b.subscribers[topic] = append(b.subscribers[topic], subscriber)
	defer b.mu.Unlock()

	return subscriber
}


func (b *Broker) Unsubscribe(topic string, subscriber *Subscriber) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if subscribers, found := b.subscribers[topic]; found {
	for i, sub := range subscribers {
		if sub == subscriber {
			b.subscribers[topic] = append(subscribers[:i], subscribers[i+1:]...)
			close(subscriber.Channel)      // Close the channel to signal completion
			close(subscriber.Unsubscribe)   // Close the unsubscribe channel
			return
		}
	}
   }
}

func (b *Broker) Publish(topic string, content interface{}) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if subscribers, found := b.subscribers[topic]; found {
        
		for _, subscriber := range subscribers {
			subscriber.Channel <- Message{Topic: topic, Content: content}
		}
	} else {
        fmt.Println("No subscriber found for topic");
    }
}


func main() {
	broker := NewBroker()

	subscriber := broker.Subscribe("topic_1")
	go func() {
		for {
			select {
			case msg, ok := <-subscriber.Channel:
				if !ok {
					fmt.Println("Subscriber channel closed.")
					return
				}
				fmt.Printf("Received: %v\n", msg.Content)
			case <-subscriber.Unsubscribe:
				fmt.Println("Unsubscribed.")
				return
			}
		}
	}()

	broker.Publish("topic_1", "my name is!")
	broker.Publish("topic_1", "This is a test message.")

	time.Sleep(2 * time.Second)
	broker.Unsubscribe("topic_1", subscriber)
    
	broker.Publish("topic_1", "This message won't be received.")

	time.Sleep(time.Second)
}