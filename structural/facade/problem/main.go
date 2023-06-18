package main

import (
	"errors"
	"log"
)

type Product struct {
	Name  string
	Price float32
}

type Inventory struct {
	products []Product
}

func (iv Inventory) Lookup(name string) (*Product, error) {
	for _, product := range iv.products {
		if product.Name == name {
			return &product, nil
		}
	}

	return nil, errors.New("product not found")
}

type Account struct {
	Name    string
	balance float32
}

func (acc *Account) Deposit(money float32)  { acc.balance += money }
func (acc *Account) Withdraw(money float32) { acc.balance -= money }
func (acc *Account) GetBalance() float32    { return acc.balance }

type AccountStorage struct {
	accounts []Account
}

func (as AccountStorage) Lookup(name string) (*Account, error) {
	for _, acc := range as.accounts {
		if acc.Name == name {
			return &acc, nil
		}
	}

	return nil, errors.New("account not found")
}

var inventory = Inventory{
	products: []Product{
		{Name: "Apple", Price: 2.5},
		{Name: "Orange", Price: 3.0},
	},
}

var accountStorage = AccountStorage{
	accounts: []Account{
		{Name: "VIP", balance: 1000},
		{Name: "Economic", balance: 300},
	},
}

func main() {
	// Case 1: Buy a product with an account
	productName := "Apple"
	accountName := "VIP"

	product, err := inventory.Lookup(productName)

	if err != nil {
		log.Fatal(err)
	}

	account, err := accountStorage.Lookup(accountName)

	if err != nil {
		log.Fatal(err)
	}

	if account.GetBalance() < product.Price {
		log.Fatal("not enough balance")
	}

	account.Withdraw(product.Price)
	// And more step to finish buying process....

	// Problem: I have to do it myself, too many steps to take. I'm
	// not sure if I do correctly
}
