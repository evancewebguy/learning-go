package main

import "fmt"

// Declaration: Functions are declared using the func keyword, followed by the function name, parameters, return types, and the function body.
func add(a, b int) int {
	return a + b
}

func compute(a, b int) (int, int) {
	sum := a + b
	diff := a - b

	return sum, diff
}

// Variadic functions: these are functions that can accept a variable number of arguments using ellipsis (`...`)
func sum(numbers ...int) int {
	total := 0

	for _, num := range numbers {
		total += num
	}
	return total
}

// closures
func outer() func() {
	count := 0

	//Inner function (closure) has access to 'count'
	increment := func() {
		count++
		fmt.Println(count)
	}
	return increment
}

func multiplier(factor int) func(int) int {
	// 'factor' is captured by the closure
	return func(x int) int {
		return x * factor
	}
}

// Functions as a type
func operation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

//end of closure

// Defer: schedule a function call to run after the function completes.
func testDefer() {
	defer fmt.Println("Defer: - This will be executed last")
	fmt.Println("without defer: - This will be executed first")
}

// Panic: Interrupt normal execution when something unexpected happens.
func testPanic() {
	panic("Panic - Something went wrong")
}

// Recover: Used to gain control after a panic.
func testRecover() {
	defer func() {
		if z := recover(); z != nil {
			fmt.Println("Recovered from panic:", z)
		}
	}()
	panic("Recover: - something went wrong in this panic")
}

func main() {
	result := add(28, 5)

	fmt.Println(result)
	fmt.Println("------")

	total, difference := compute(8, 3)

	fmt.Println(total)
	fmt.Println(difference)
	fmt.Println("------")

	totalSum := sum(1, 2, 3, 4)
	fmt.Println(totalSum)

	fmt.Println("---Anonymous function assigned to a variable---")

	// Anonymous function assigned to a variable
	addNum := func(a, b int) int {
		return a + b
	}

	resultAdd := addNum(3, 4)
	fmt.Println(resultAdd)

	fmt.Println("---Closures---")
	//example 1
	counter := outer()

	//Each call to 'counter' maintains its own 'count'
	counter() //output: 1
	counter() //output: 2
	counter() //output: 3

	//example 2
	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println(double(5))
	fmt.Println(triple(5))

	fmt.Println("---Functions as a type-----")

	resultO := operation(3, 4, add)
	fmt.Println(resultO)

	fmt.Println("----Defer----")
	testDefer()

	fmt.Println("----Panic----")
	testPanic()

	fmt.Println("----Recover:----")
	testRecover()
}
