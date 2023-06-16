package main

import "fmt"

type WebPage interface {
	RenderHeader()
	RenderBody()
	RenderFooter()
	Render()
}

type HomePage struct{}

func (HomePage) RenderHeader() { fmt.Println("This is header") }
func (HomePage) RenderBody()   { fmt.Println("This is body of home page") }
func (HomePage) RenderFooter() { fmt.Println("This is footer") }
func (p HomePage) Render() {
	p.RenderHeader() // can be the same for other pages
	p.RenderBody()   // different part
	p.RenderFooter() // can be the same for other pages
}
