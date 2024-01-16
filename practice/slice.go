package main

import "fmt"

func main() {
	c := make([]string, 5)
	fmt.Println("test:", c)

	c[0] = "a"
	c[1] = "b"
	c[2] = "c"
	c[3] = "d"
	c[4] = "e"
	fmt.Println("set:", c)

	l := c[3:5]
	fmt.Println("sl1:", l)

	m := c[3:]
	fmt.Println("slm:", m)
}
