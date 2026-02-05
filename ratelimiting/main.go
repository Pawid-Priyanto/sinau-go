package main

import (
	"fmt"
	"time"
)

func main() {

	request := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		request <- i
	}
	close(request)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range request {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for range 3 {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequest := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequest <- i
	}

	close(burstyRequest)

	for req := range burstyRequest {
		<-burstyLimiter
		fmt.Println("requset", req, time.Now())
	}
}
