package main

import "fmt"

type area struct {
	length int
	width  int
}

func (a *area) area() int {
	return a.length * a.width
}

func (a area) area1() int {
	return a.length * a.width
}

func area2(length, width int) int {
	return length * width
}

func main() {
	ar := area{5, 5}

	fmt.Println(ar.area())

	fmt.Println(ar.area1())

	fmt.Println(area2(6, 6))
}
