package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {

		<-timer2.C
		fmt.Println("Timer 2 Fired")
	}()

	stop := timer2.Stop()

	if stop {
		fmt.Println("timer 2 stop")
	}

	time.Sleep(2 * time.Second)
}
