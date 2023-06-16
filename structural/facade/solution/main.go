package main

import (
	"errors"
	"fmt"
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
	accounts []*Account
}

func (as AccountStorage) Lookup(name string) (*Account, error) {
	for _, acc := range as.accounts {
		if acc.Name == name {
			return acc, nil
		}
	}

	return nil, errors.New("account not found")
}

type FacadeService struct {
	inventory      Inventory
	accountStorage AccountStorage
}

func (s *FacadeService) BuyProduct(name, accountName string) error {
	product, err := s.inventory.Lookup(name)

	if err != nil {
		return err
	}

	account, err := s.accountStorage.Lookup(accountName)

	if err != nil {
		return err
	}

	if account.GetBalance() < product.Price {
		return errors.New("not enough balance in account")
	}

	account.Withdraw(product.Price)
	// And more step to finish buying process....

	return nil
}

func (s *FacadeService) Deposit(accountName string, money float32) error {
	account, err := s.accountStorage.Lookup(accountName)

	if err != nil {
		return err
	}

	account.Deposit(money)

	return nil
}

func (s *FacadeService) FetchBalance(accountName string) float32 {
	account, err := s.accountStorage.Lookup(accountName)

	if err != nil {
		return 0
	}

	return account.GetBalance()
}

func NewFacadeService() FacadeService {
	return FacadeService{
		inventory: Inventory{
			products: []Product{
				{Name: "Apple", Price: 2.5},
				{Name: "Orange", Price: 3.0},
			},
		},
		accountStorage: AccountStorage{
			accounts: []*Account{
				{Name: "VIP", balance: 1000},
				{Name: "Economic", balance: 300},
			},
		},
	}
}

func main() {
	service := NewFacadeService()

	// Case 1: Buy a product with an account
	productName := "Apple"
	accountName := "VIP"

	if err := service.BuyProduct(productName, accountName); err != nil {
		log.Fatal(err)
	}

	// Check my balance
	fmt.Println("Account Balance:", service.FetchBalance(accountName))

	// Case 2: Deposit 100 into VIP Account

	if err := service.Deposit(accountName, 100); err != nil {
		log.Fatal(err)
	}

	// Check my balance again
	fmt.Println("Account Balance:", service.FetchBalance(accountName))
}
