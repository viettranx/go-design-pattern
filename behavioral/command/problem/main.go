package main

import "fmt"

type Value struct {
	val int
}

func (v *Value) Add(n int) { v.val += n }
func (v *Value) Sub(n int) { v.val -= n }
func (v *Value) Val() int  { return v.val }

func NewValue(v int) Value {
	return Value{v}
}

func main() {
	v := NewValue(10)

	v.Add(5)
	v.Add(7)
	v.Sub(2)
	v.Sub(3)
	v.Add(4)

	fmt.Println(v.Val())

	// how can I undo it?
}
