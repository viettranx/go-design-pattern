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

func CreateNotifier(t string) Notifier {
	if t == "sms" {
		return SMSNotifier{}
	}

	return EmailNotifier{}
}

func main() {
	s := Service{
		notifier: CreateNotifier("email"),
	}

	s.SendNotification("Hello world")
}
