package main

import (
	"fmt"
	"sync"
	"time"
)

var sharedResource int

// var mu sync.Mutex
var rwmu sync.RWMutex

//func increament() {
//	mu.Lock()
//	defer mu.Unlock()
//	sharedResource++
//}
//
//func decreament() {
//	mu.Lock()
//	defer mu.Unlock()
//	sharedResource--
//}

func read() {
	rwmu.RLock()
	defer rwmu.RUnlock()
	fmt.Println("Read:", sharedResource)

}
func write() {
	rwmu.Lock()
	defer rwmu.Unlock()

	sharedResource++
	fmt.Println("Write:", sharedResource)
}

func main() {
	//Example two

	//TOPIC: RWMutex
	var wg sync.WaitGroup

	//readers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			read()
		}()
	}

	//write
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			write()
		}()
	}
	wg.Wait()
	fmt.Println("Final value of sharedResource:", sharedResource)

	//TOPIC: Mutex
	//var wg sync.WaitGroup
	//
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		increament()
	//	}()
	//}
	//wg.Wait()
	//
	//fmt.Println("Final value of sharedResource:", sharedResource)

	//Example one
	//now := time.Now()
	//userID := 10
	//
	//respch := make(chan string)
	//
	//go fetchUserData(userID, respch)
	//go fetchUserRecommendation(userID, respch)
	//go fetchUserLikes(userID, respch)
	//
	//close(respch)
	//for resp := range respch {
	//	fmt.Println(resp)
	//}
	//
	//fmt.Println(time.Since(now))

}

func fetchUserData(userID int, respch chan string) {
	time.Sleep(80 * time.Millisecond)
	respch <- "user data"
}

func fetchUserRecommendation(userID int, respch chan string) {
	time.Sleep(120 * time.Millisecond)
	respch <- "user Recommendations"
}
func fetchUserLikes(userID int, respch chan string) {
	time.Sleep(50 * time.Millisecond)
	respch <- "user Likes"
}
