package main

import "go-design-pattern/creational/builder/solution/option-function/internal"

func main() {
	service := internal.NewService(
		internal.WithName("Complex Service"),
		internal.WithStdLogger(),
		internal.WithEmailNotifier(),
		internal.WithMySQLDataLayer(),
	)

	service.DoBusiness()
}
