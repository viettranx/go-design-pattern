package main

import (
	"fmt"
	"time"
)

type DataStorage interface {
	GetValue() int
}

type RealDataStorage struct{}

func (RealDataStorage) GetValue() int {
	time.Sleep(time.Second * 2)

	return 100
}

type ValueService struct {
	storage DataStorage
}

func (s ValueService) FetchValue() int {
	return s.storage.GetValue()
}

func main() {
	value := ValueService{storage: RealDataStorage{}}.FetchValue()
	// It's too low...
	fmt.Println(value)

	// I would like to use some cache to speed up, but where I can put more logic for it?
}
