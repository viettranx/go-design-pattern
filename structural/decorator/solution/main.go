package main

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}

func (EmailNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (Sender: Email)\n", message)
}

type SMSNotifier struct{}

func (SMSNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (Sender: SMS)\n", message)
}

type TelegramNotifier struct{}

func (notifier TelegramNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (Sender: Telegram)\n", message)
}

type NotifierDecorator struct {
	core     *NotifierDecorator
	notifier Notifier
}

func (nd NotifierDecorator) Send(message string) {
	nd.notifier.Send(message)

	if nd.core != nil {
		nd.core.Send(message)
	}
}

// Like add to stack

func (nd NotifierDecorator) Decorate(notifier Notifier) NotifierDecorator {
	return NotifierDecorator{
		core:     &nd,
		notifier: notifier,
	}
}

func NewNotifierDecorator(notifier Notifier) NotifierDecorator {
	return NotifierDecorator{notifier: notifier}
}

type Service struct {
	notifier Notifier
}

func (s Service) SendNotification(message string) {
	s.notifier.Send(message)
}

func main() {
	notifier := NewNotifierDecorator(EmailNotifier{}).
		Decorate(SMSNotifier{}).
		Decorate(TelegramNotifier{})

	s := Service{
		notifier: notifier,
	}

	s.SendNotification("Hello world")
}
