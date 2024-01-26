package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime()
}

func timeformatter(currentTime time.Time) func() string {
	formatToReturn := "2006-01-02 15:04:05"
	return func() string {
		return currentTime.Format(formatToReturn)
	}
}

// current time
func currentTime() {
	currentTime := time.Now()
	fmt.Println("current Time:", currentTime)
	fmt.Println("----------------")

	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	fmt.Println("formatted Time:", formattedTime)
	fmt.Println("----------------")

	// Call the function returned by timeformatter to get the formatted time
	formattedTime2 := timeformatter(currentTime)()
	fmt.Println("formatted Time:", formattedTime2)
	fmt.Println("----------------")

	futureTime := currentTime.Add(24 * time.Hour)
	pastTime := currentTime.Add(-2 * time.Hour)

	fmt.Println("Future time:", futureTime)
	fmt.Println("past Time :", pastTime)

}
