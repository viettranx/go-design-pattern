package main

import (
	"log"
	"sync"
	"time"
)

type config struct {
	logAllowed bool
}

func (c config) LogAllowed() bool { return c.logAllowed }

func NewConfig(allowed bool) config {
	return config{logAllowed: allowed}
}

// This is all about Singleton Pattern
// the rest is demo for data racing in concurrency Golang
var singletonApp = &application{once: sync.Once{}}

func GetApplication() *application {
	return singletonApp
}

type application struct {
	once sync.Once
	cfg  *config
}

func (app *application) GetConfig() *config {
	//if app.cfg == nil {
	//	log.Println("It should be run only once. But you'll it many times!!")
	//	app.loadConfig()
	//}
	//
	//return app.cfg

	if app.cfg == nil {
		app.once.Do(func() {
			log.Println("Loading config once and forever.")
			app.loadConfig()
		})
	}

	return app.cfg
}

func (app *application) loadConfig() {
	time.Sleep(100) // demo delay time of loading config
	app.cfg = &config{logAllowed: true}
}

func main() {
	// Demo 1000 requests to service at a same time (1000 RPS)
	// I made this code for simple demo, not a real practice!

	rps := 1000
	wg := sync.WaitGroup{}
	wg.Add(rps)

	for i := 1; i <= rps; i++ {
		go func(idx int) {
			requestHandler(idx)

			if GetApplication().GetConfig().LogAllowed() {
				log.Printf("Request %d handled successfully.\n", idx)
			}

			wg.Done()
		}(i)
	}

	wg.Wait()
}

func requestHandler(requestIdx int) {
	// YES. Now I can get the config

	if GetApplication().GetConfig().LogAllowed() {
		log.Printf("Handling request %d... please wait.\n", requestIdx)
	}
}
