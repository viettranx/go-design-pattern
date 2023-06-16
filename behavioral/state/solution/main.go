package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrInvalidAction = errors.New("invalid action")

type OrderState interface {
	Cancel() error
	Pay() error
	Deliver() error
	Finish() error
	String() string
}

type OrderStateCancelled struct {
	order *Order
}

func (OrderStateCancelled) Cancel() error  { return ErrInvalidAction }
func (OrderStateCancelled) Pay() error     { return ErrInvalidAction }
func (OrderStateCancelled) Deliver() error { return ErrInvalidAction }
func (OrderStateCancelled) Finish() error  { return ErrInvalidAction }
func (OrderStateCancelled) String() string { return "cancelled" }

type OrderStateFinished struct {
	order *Order
}

func (OrderStateFinished) Cancel() error  { return ErrInvalidAction }
func (OrderStateFinished) Pay() error     { return ErrInvalidAction }
func (OrderStateFinished) Deliver() error { return ErrInvalidAction }
func (OrderStateFinished) Finish() error  { return ErrInvalidAction }
func (OrderStateFinished) String() string { return "finished" }

type OrderStateCreated struct {
	order *Order
}

func (state OrderStateCreated) Cancel() error {
	state.order.updateState(OrderStateCancelled{order: state.order})
	return nil
}

func (state OrderStateCreated) Pay() error {
	state.order.updateState(OrderStatePaid{order: state.order})
	return nil
}

func (OrderStateCreated) Deliver() error { return ErrInvalidAction }
func (OrderStateCreated) Finish() error  { return ErrInvalidAction }
func (OrderStateCreated) String() string { return "created" }

type OrderStatePaid struct {
	order *Order
}

func (OrderStatePaid) Cancel() error { return ErrInvalidAction }
func (OrderStatePaid) Pay() error    { return ErrInvalidAction }
func (state OrderStatePaid) Deliver() error {
	state.order.updateState(OrderStateDelivered{order: state.order})
	return nil
}
func (OrderStatePaid) Finish() error  { return ErrInvalidAction }
func (OrderStatePaid) String() string { return "paid" }

type OrderStateDelivered struct {
	order *Order
}

func (OrderStateDelivered) Cancel() error  { return ErrInvalidAction }
func (OrderStateDelivered) Pay() error     { return ErrInvalidAction }
func (OrderStateDelivered) Deliver() error { return ErrInvalidAction }
func (state OrderStateDelivered) Finish() error {
	state.order.updateState(OrderStateFinished{order: state.order})
	return nil
}
func (OrderStateDelivered) String() string { return "delivered" }

type Order struct {
	state OrderState
}

func NewOrder() *Order {
	order := &Order{}
	initState := OrderStateCreated{order: order}
	order.state = initState

	return order
}

func (o *Order) updateState(state OrderState) {
	o.state = state
	log.Printf("order has changed state to: %s\n", state)
}

func (o *Order) CurrentState() OrderState { return o.state }

func (o *Order) Cancel() error { return o.state.Cancel() }

func (o *Order) Pay() error { return o.state.Pay() }

func (o *Order) Deliver() error { return o.state.Deliver() }

func (o *Order) Finish() error { return o.state.Finish() }

func main() {
	order := NewOrder()

	fmt.Println("Order state:", order.CurrentState())

	// Invalid action
	if err := order.Finish(); err != nil {
		log.Println(err)
	}

	// OK
	if err := order.Pay(); err != nil {
		log.Println(err)
	}

	// OK
	if err := order.Deliver(); err != nil {
		log.Println(err)
	}

	// OK
	if err := order.Finish(); err != nil {
		log.Println(err)
	}
}

// Another simple example

// Green --> Yellow --> Red --> back to Green

type TrafficLightState interface {
	Next()
	Render()
}

type TrafficLight struct {
	state TrafficLightState
}

func (l *TrafficLight) changeState(state TrafficLightState) { l.state = state }
func (l *TrafficLight) Render()                             { l.state.Render() }

type TrafficLightGreen struct{ light *TrafficLight }
type TrafficLightYellow struct{ light *TrafficLight }
type TrafficLightRed struct{ light *TrafficLight }

func (l TrafficLightGreen) Next()   { l.light.changeState(TrafficLightYellow{light: l.light}) }
func (l TrafficLightGreen) Render() { fmt.Println("Traffic light: green on") }

func (l TrafficLightYellow) Next()   { l.light.changeState(TrafficLightRed{light: l.light}) }
func (l TrafficLightYellow) Render() { fmt.Println("Traffic light: yellow on") }

func (l TrafficLightRed) Next()   { l.light.changeState(TrafficLightGreen{light: l.light}) }
func (l TrafficLightRed) Render() { fmt.Println("Traffic light: red on") }
