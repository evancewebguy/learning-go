package main

import "fmt"

func main() {
	//Pointers

	//example 2
	var x int = 42
	var ptr *int

	ptr = &x // Assign the address of x to ptr

	// Dereference ptr to get the value stored at the address it points to
	fmt.Println(ptr)
	fmt.Println(*ptr)

	//exapmple 1
	// Declare a variable
	//var x int = 42
	//
	//// Declare a pointer variable and assign the address of x to it
	//var ptr *int = &x
	//
	//// Print the value and address of x
	//fmt.Println("Value of x:", x)
	//fmt.Println("Address of x:", &x)
	//
	//// Print the value and address stored in the pointer variable
	//fmt.Println("Value stored in ptr:", *ptr)
	//fmt.Println("Address stored in ptr:", ptr)
	//
	//// Modify the value of x through the pointer
	//*ptr = 99
	//
	//// Print the modified value of x
	//fmt.Println("Modified value of x:", x)

}
