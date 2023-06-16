package main

import "fmt"

type Item interface {
	Cost() float32
}

type RealItem struct {
	Name  string
	Price float32
}

func (item RealItem) Cost() float32 {
	return item.Price
}

type Box struct {
	children []Item
}

func (b Box) Cost() float32 {
	var cost float32 = 0.0

	for _, item := range b.children {
		cost += item.Cost()
	}

	return cost
}

func CreatePackage() Item {
	return Box{
		children: []Item{
			RealItem{
				Name:  "Mouse",
				Price: 20.5,
			},
			Box{
				children: []Item{
					RealItem{
						Name:  "Keyboard",
						Price: 60,
					},
					RealItem{
						Name:  "Charger",
						Price: 15,
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
