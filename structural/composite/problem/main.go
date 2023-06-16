package main

import "fmt"

type Item struct {
	Name  string
	Price float32
	// Ops, an item always has children!! Why??
	children []Item
}

func (item Item) Cost() float32 {
	cost := item.Price

	for _, child := range item.children {
		cost += child.Cost()
	}

	return cost
}

func CreatePackage() Item {
	return Item{
		// It's a box, not an item
		Name: "root box",
		// so it have no price
		Price: 0,
		// but it has children
		children: []Item{
			// Here it is!! My real item here.
			{
				Name:     "Mouse",
				Price:    20.5,
				children: nil,
			},
			// Another box contains items
			{
				Name:  "sub box",
				Price: 0,
				children: []Item{
					{
						Name:     "Keyboard",
						Price:    60,
						children: nil,
					},
					{
						Name:     "Charger",
						Price:    15,
						children: nil,
					},
				},
			},
		},
	}
}

func main() {
	myPackage := CreatePackage()

	fmt.Println(myPackage.Cost())
}
