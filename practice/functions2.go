package main

import "fmt"

///LEANING fUNCTION CLOSURES

func simpleclosure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// CAPTURING VARIABLES
func xmultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func main() {
	// example two
	double := xmultiplier(2)
	fmt.Println(double(3))
	fmt.Println("----------------")

	///example one
	counter := simpleclosure()

	fmt.Println(counter())
	fmt.Println(counter())
}
