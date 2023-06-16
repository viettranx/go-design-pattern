package main

import "fmt"

type Point struct {
	X, Y int
}

type Node struct {
	Value    Point
	Children []Point
}

func main() {
	p1 := Point{1, 2}
	p2 := p1

	p1.X = 3
	fmt.Println(p1, p2) // p2 copy from p1

	n1 := Node{
		Value: p1,
		Children: []Point{
			{1, 2},
			{2, 3},
		},
	}

	n2 := n1

	n1.Children[0].X = 10

	fmt.Println(n1, n2) // n2 has `children` reference to n1. So n2 is `shallow copy`.
}
