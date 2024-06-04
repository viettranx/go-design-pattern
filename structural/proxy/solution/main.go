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

type ProxyDataStorage struct {
	cachedValue *int
	realStorage DataStorage
}

func (s *ProxyDataStorage) GetValue() int {
	// I would like to ignore data racing issue here to keep it simple
	if val := s.cachedValue; val != nil {
		return *val
	}

	val := s.realStorage.GetValue()
	s.cachedValue = &val

	return val
}

func NewProxyDataStorage(realStorage RealDataStorage) DataStorage {
	return &ProxyDataStorage{realStorage: realStorage}
}

type ValueService struct {
	storage DataStorage
}

func (s ValueService) FetchValue() int {
	return s.storage.GetValue()
}

func main() {
	vs := ValueService{
		storage: NewProxyDataStorage(RealDataStorage{}),
	}

	// It's too low at the first time
	fmt.Println(vs.FetchValue())

	// Now it return instantly because of caching (proxy layer)
	fmt.Println(vs.FetchValue())
}
