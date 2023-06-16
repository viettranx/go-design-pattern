package main

import "go-design-pattern/creational/builder/solution/builder/internal"

func main() {
	director := internal.NewDirector()
	builder := internal.NewBuilder()

	service := director.BuildService(builder)
	service.DoBusiness()
}
