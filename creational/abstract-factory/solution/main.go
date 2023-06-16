package main

import (
	"errors"
	"fmt"
	"log"
)

type Drink interface {
	Drink()
}
type Food interface {
	Eat()
}

type Voucher struct {
	Drink
	Food
}

type Coffee struct{}

func (Coffee) Drink() {
	fmt.Println("It's coffee, drinkable")
}

type Beer struct{}

func (Beer) Drink() {
	fmt.Println("It's beer, drinkable")
}

type Cake struct{}

func (Cake) Eat() {
	fmt.Println("It's cake, eatable")
}

type GrilledOctopus struct{}

func (GrilledOctopus) Eat() {
	fmt.Println("It's Grilled Octopus, eatable")
}

type VoucherAbstractFactory interface {
	GetDrink() Drink
	GetFood() Food
}

type CoffeeMorningVoucherFactory struct{}

func (CoffeeMorningVoucherFactory) GetDrink() Drink { return Coffee{} }
func (CoffeeMorningVoucherFactory) GetFood() Food   { return Cake{} }

type DrinkEveningVoucherFactory struct{}

func (DrinkEveningVoucherFactory) GetDrink() Drink { return Beer{} }
func (DrinkEveningVoucherFactory) GetFood() Food   { return GrilledOctopus{} }

func GetVoucherFactory(campaignName string) (VoucherAbstractFactory, error) {
	if campaignName == "creative-morning" {
		return CoffeeMorningVoucherFactory{}, nil
	}

	if campaignName == "chill-all-night-long" {
		return DrinkEveningVoucherFactory{}, nil
	}

	return nil, errors.New("campaign not found")
}

func GetVoucher(factory VoucherAbstractFactory) Voucher {
	return Voucher{
		Drink: factory.GetDrink(),
		Food:  factory.GetFood(),
	}
}

func main() {
	voucherFactory, err := GetVoucherFactory("creative-morning")

	if err != nil {
		log.Fatal(err)
	}

	myVoucher := GetVoucher(voucherFactory)

	fmt.Println("I'm happy with this voucher and come back to use it next time.", myVoucher)
}
