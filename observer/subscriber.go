package main

type Subscriber interface {
	Id() string
	React(msg string)
}

type subscriber struct {
	subId string
}

func NewSubscriber(subId string) Subscriber {
	return &subscriber{subId: subId}
}

func (s *subscriber) Id() string {
	return s.subId
}

func (s *subscriber) React(msg string) {
	println(s.subId + " received: " + msg)
}
