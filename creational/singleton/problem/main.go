package main

import (
	"log"
	"sync"
)

type config struct {
	logAllowed bool
}

func (c config) LogAllowed() bool { return c.logAllowed }

func NewConfig(allowed bool) config {
	return config{logAllowed: allowed}
}

func main() {
	// Demo 1000 requests to service at a same time (1000 RPS)
	// I made this code for simple demo, not a real practice!

	rps := 100
	wg := sync.WaitGroup{}
	wg.Add(rps)

	config := NewConfig(true)

	for i := 1; i <= rps; i++ {
		go func(idx int) {
			requestHandler(idx)

			if config.LogAllowed() {
				log.Printf("Request %d handled successfully.\n", idx)
			}

			wg.Done()
		}(i)
	}

	wg.Wait()
}

func requestHandler(requestIdx int) {
	// I have some log to print here
	// I have to know the config that if it was allowed print log.
	// but I cannot modify definition of this method.
}
