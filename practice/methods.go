package main

import "fmt"

type rect struct {
	width  int
	height int
}

// Pointer Receivers
// the method operates receiver pointer on the actual value, allowing modification.
// use it when the method needs to modify the receiver value.
// suitable for mutual operations on the receiver
func (r *rect) area() int {
	return r.width * r.height
}

// Value Receivers:
// it operates in a copy of the value its called on.
// doesnt modify the original value.
// mostly operate on a read only operations on the receiver.
// Efficient for small and simple types
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {

	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
