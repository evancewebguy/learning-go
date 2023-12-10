package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {
	colors := []string{"white", "red", "green"}
	fmt.Println(colors)

	// Math

	circleRadius := 15.5
	circumference := math.Pi * circleRadius * 2
	fmt.Printf("circumference %.2f\n", circumference)

	// Date and time

	now := time.Now()
	fmt.Println("This is the now time at:", now)

	d := time.Date(2023, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("This time is at:", d)
	fmt.Println(d.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Fri Nov 10 23:00:00 2023")
	fmt.Printf("The time of the parsedTime is %T\n", parsedTime)

	// Convert string to float number
	str := "40.98"
	fmt.Println(strconv.ParseFloat(str, 32))
	fmt.Println(strconv.ParseFloat(str, 64))

	//Pointers

	anInt := 42
	var p = &anInt
	fmt.Println("The value of p:", *p)
}
