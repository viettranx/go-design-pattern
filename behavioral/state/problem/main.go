package main

import "errors"

type OrderState int

// Finite States
const (
	OrderStateCreated = iota
	OrderStateCancelled
	OrderStatePaid
	OrderStateDelivered
	OrderStateFinished
)

//``` [Start] --> Created --> Cancelled --> [End]
//					|
//					|
//					Paid  --> Delivered --> Finished --> [End]
//```

var ErrInvalidAction = errors.New("invalid action")

type Order struct {
	state OrderState
}

func NewOrder() Order {
	return Order{state: OrderStateCreated}
}

func (o *Order) Cancel() error {
	if o.state != OrderStateCreated {
		return ErrInvalidAction
	}

	o.state = OrderStateCancelled
	return nil
}

func (o *Order) Pay() error {
	if o.state != OrderStateCreated {
		return ErrInvalidAction
	}

	o.state = OrderStatePaid
	return nil
}

func (o *Order) Deliver() error {
	if o.state != OrderStatePaid {
		return ErrInvalidAction
	}

	o.state = OrderStateDelivered
	return nil
}

func (o *Order) Finish() error {
	if o.state != OrderStateDelivered {
		return ErrInvalidAction
	}

	o.state = OrderStateFinished
	return nil
}
