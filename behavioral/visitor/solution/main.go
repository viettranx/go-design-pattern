package main

import "fmt"

type Visitor interface {
	// Don't do this, we have to type check for every parameter
	// It's god-function (too complex)
	//Visit(any)

	VisitProduct(product Product)
	VisitCategory(category Category)
	VisitUser(user User)

	// Just feel free for add more in future
}

type Visitable interface {
	Accept(v Visitor)
}

type PrinterVisitor struct{}

func (PrinterVisitor) VisitProduct(p Product)   { fmt.Printf("product: %s\n", p.Name) }
func (PrinterVisitor) VisitCategory(c Category) { fmt.Printf("category: %s\n", c.Title) }
func (PrinterVisitor) VisitUser(u User)         { fmt.Printf("user: %s %s\n", u.FirstName, u.LastName) }

type Category struct {
	Title string
}
type Product struct {
	Name string
}
type User struct {
	FirstName string
	LastName  string
}

func (c Category) Accept(v Visitor) { v.VisitCategory(c) }
func (p Product) Accept(v Visitor)  { v.VisitProduct(p) }
func (u User) Accept(v Visitor)     { v.VisitUser(u) }

func main() {
	items := []Visitable{
		Category{Title: "Programming"},
		Product{Name: "Visitor Design Pattern"},
		User{FirstName: "Viet", LastName: "Tran"},
	}

	printerVisitor := PrinterVisitor{}
	jsonEncoderVisitor := JSONEncoderVisitor{}

	for i := range items {
		items[i].Accept(printerVisitor)
		items[i].Accept(jsonEncoderVisitor)
	}
}

type JSONEncoderVisitor struct{}

func (JSONEncoderVisitor) VisitProduct(p Product)   { fmt.Printf("{\"name\": \"%s\"}", p.Name) }
func (JSONEncoderVisitor) VisitCategory(c Category) { fmt.Printf("{\"title\": \"%s\"}", c.Title) }
func (JSONEncoderVisitor) VisitUser(u User) {
	fmt.Printf("{\"first_name\": \"%s\", \"last_name\": \"%s\"}", u.FirstName, u.LastName)
}
