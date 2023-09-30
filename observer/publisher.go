package main

type Publisher interface {
	Subscribe(sub Subscriber)
	Unsubscribe(id string)
	Emit(msg string)
}

type publisher struct {
	subscribers map[string]Subscriber
}

func NewPublisher() Publisher {
	return &publisher{
		subscribers: make(map[string]Subscriber),
	}
}

func (p *publisher) Subscribe(sub Subscriber) {
	p.subscribers[sub.Id()] = sub
}

func (p *publisher) Unsubscribe(id string) {
	delete(p.subscribers, id)
}

func (p *publisher) Emit(msg string) {
	for _, sub := range p.subscribers {
		sub.React(msg)
	}
}
