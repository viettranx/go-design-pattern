package main

import (
	"fmt"
)

type WebPage interface {
	RenderHeader()
	RenderBody()
	RenderFooter()
}

type WebTemplate struct{}

func (WebTemplate) RenderHeader() { fmt.Println("This is header") }
func (WebTemplate) RenderFooter() { fmt.Println("This is footer") }

type HomePage struct {
	WebTemplate // embed
}

func (HomePage) RenderBody() { fmt.Println("This is body of home page") }

type ProductPage struct {
	WebTemplate // embed
}

func (ProductPage) RenderBody() { fmt.Println("This is body of product page") }

func renderWebsite(website WebPage) {
	website.RenderHeader()
	website.RenderBody()
	website.RenderFooter()
}

func main() {
	homePage := HomePage{}
	productPage := ProductPage{}

	renderWebsite(homePage)
	renderWebsite(productPage)
}

// Another example

func mapInt(numbers []int, f func(int) int) []int {
	result := make([]int, len(numbers))

	for i := range numbers {
		result[i] = f(numbers[i])
	}

	return result
}
