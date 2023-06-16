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

type NotificationService struct {
	notifier Notifier
}

func (s NotificationService) SendNotification(message string) {
	s.notifier.Send(message)
}

func main() {
	s := NotificationService{
		notifier: EmailNotifier{},
	}

	s.SendNotification("Hello world")
}
