package main

import "fmt"

type Printer interface {
	Print() string
}

type Category struct{}
type Product struct{}
type User struct{}

func (Category) Print() string { return "printing category" }
func (Product) Print() string  { return "printing product" }
func (User) Print() string     { return "printing user" }

func doPrint(p Printer) {
	fmt.Println(p.Print())
}

// Now I have a new interface,
// so I have to modify Category, Product, User to implement it
// But Category, Product and User said: No, you don't. Because it's not
// our business

type Encoder interface {
	Encode() string
}

// Another method for executing Encoder
func doEncode(encoder Encoder) {
	fmt.Println(encoder.Encode())
}

func main() {
	items := []Printer{Category{}, Product{}, User{}}

	for i := range items {
		doPrint(items[i])
	}
}
