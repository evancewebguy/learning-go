package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//EXAMPLE FIVE
	//var wg sync.WaitGroup
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	fmt.Println("1st gorotine sleeping")
	//	time.Sleep(1)
	//}()
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	fmt.Println("2nd gorotine running")
	//	time.Sleep(2)
	//}()
	//wg.Wait()
	//fmt.Println("All goroutine complete")

	//EXAMPLE Four - correct but verbose
	//var wg sync.WaitGroup
	//for _, salutation := range []string{"Hello", "Greetings", "Good Day"} {
	//	wg.Add(1)
	//	go func(salutation string) {
	//		defer wg.Done()
	//		fmt.Println(salutation)
	//	}(salutation)
	//}
	//wg.Wait()

	//EXAMPLE: THREE - for wrong answ
	//var wg sync.WaitGroup
	//for _, salutation := range []string{"Hello", "Greetings", "Good Day"} {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		fmt.Println(salutation)
	//	}()
	//}
	//wg.Wait()

	//EXAMPLE: two
	//var wg sync.WaitGroup
	//salutation := "Hello"
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	salutation = "welcome"
	//}()
	//wg.Wait()
	//fmt.Println(salutation)

	//example one
	//var wg sync.WaitGroup
	//sayHello := func() {
	//	defer wg.Done()
	//	fmt.Println("hello")
	//}
	//wg.Add(1)
	//go sayHello()
	//wg.Wait()

	//TEST GO ROUTINES

	wg := &sync.WaitGroup{}
	wg.Add(3)
	wg.Wait()
	go getUsers(wg)
	go getComments(wg)
	go getLikes(wg)

	fmt.Println("everything is good")

}

func getUsers(wg *sync.WaitGroup) {
	time.Sleep(80 * time.Millisecond)
	fmt.Println("we have all users")
	defer wg.Done()
}

func getComments(wg *sync.WaitGroup) {
	time.Sleep(120 * time.Millisecond)
	fmt.Println("All commrnts are good")
	defer wg.Done()

}

func getLikes(wg *sync.WaitGroup) {
	time.Sleep(50 * time.Millisecond)
	fmt.Println("lIKES cool are good")
	defer wg.Done()
}
