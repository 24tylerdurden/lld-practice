package main

import "fmt"

type Event struct {
	id      int
	message string
}

var topic string

type Consumer interface {
	onEvent(Event)
	getId() int
}

type PubSub struct {
	consumers map[string][]Consumer
}

func NewPubSub() *PubSub {
	return &PubSub{
		consumers: make(map[string][]Consumer),
	}
}

func (p *PubSub) AddConsumer(topic string, consumer Consumer) {
	p.consumers[topic] = append(p.consumers[topic], consumer)
}

func (p *PubSub) DeleteConsumer(topic string, id int) {
	consumers := p.consumers[topic]

	filtered := []Consumer{}

	for _, consumer := range consumers {
		if consumer.getId() != id {
			filtered = append(filtered, consumer)
		}
	}

	p.consumers[topic] = filtered

}

func (p *PubSub) PublishEvent(topic string, event Event) {
	consumers := p.consumers[topic]

	if len(consumers) > 0 {
		for _, consumer := range consumers {
			consumer.onEvent(event)
		}
	}
}

type NewsConsumer struct {
	Id int
}

func NewNewsConsumer(id int) *NewsConsumer {
	return &NewsConsumer{Id: id}
}

func (n *NewsConsumer) onEvent(event Event) {
	fmt.Printf("The news consumer has received the event : %s \n", event.message)
}

func (n *NewsConsumer) getId() int {
	return n.Id
}

type SportsConsumer struct {
	Id int
}

func NewSportsConsumer(id int) *SportsConsumer {
	return &SportsConsumer{
		Id: id,
	}
}

func (s *SportsConsumer) onEvent(event Event) {
	fmt.Printf("The sports consumer has received the event : %s \n", event.message)
}

func (s *SportsConsumer) getId() int {
	return s.Id
}

// type Publisher interface {
// 	PublishEvent(topic string, event Event)
// }

// type Producer struct {
// 	publisher Publisher
// }

// func NewProducer(p Publisher) *Producer {
// 	return &Producer{
// 		publisher: p,
// 	}
// }

// func (p *Producer) Produce(topic string, event Event) {
// 	p.publisher.PublishEvent(topic, event)
// }

type Publish interface {
	PublishEvent(topic string, event Event)
}

type Producer struct {
	publisher Publish
}

func NewProducer(producer Publish) *Producer {
	return &Producer{
		publisher: producer,
	}
}

func main() {
	pubSub := NewPubSub()

	// create a publishers

	newsConsumer := NewNewsConsumer(1)
	sportsConsumer := NewSportsConsumer(1)

	pubSub.AddConsumer("news", newsConsumer)

	pubSub.AddConsumer("sports-news", sportsConsumer)

	newsEvent := Event{
		id:      1,
		message: "I am publishing a sports event",
	}

	producer := NewProducer(pubSub)

	producer.publisher.PublishEvent("sports-news", newsEvent)

}

// If I creata publishers,
// create a pubsub system
// define the publishers
//

// PubSub systems
// Publish Event
