package main

func main() {
	var p Publisher = NewPublisher()

	p.Emit("Hello, Subscribers!")

	p.Subscribe(NewSubscriber("john-doe"))

	p.Emit("message 1")

	p.Subscribe(NewSubscriber("jane-doe"))

	p.Emit("message 2")

	p.Unsubscribe("jane-doe")

	p.Emit("message 3")
}
