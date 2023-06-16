package main

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}

func (EmailNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (Sender: Email)", message)
}

type SMSNotifier struct{}

func (SMSNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (Sender: SMS)", message)
}

type Service struct {
	notifier Notifier
}

func (s Service) SendNotification(message string) {
	s.notifier.Send(message)
}

func main() {
	s := Service{
		// I don't want my users init a new notifier like this.
		// They should call to something to produce a notifier with its specific type
		// CreateNotifier(type) Notifier
		notifier: EmailNotifier{},
	}

	s.SendNotification("Hello world")
}
