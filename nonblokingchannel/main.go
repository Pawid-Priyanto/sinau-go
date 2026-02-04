package main

import "fmt"

func main() {
	messagges := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messagges:
		fmt.Println("received messages", msg)
	default:
		fmt.Println("no messages received")
	}

	msg := "hi"
	select {
	case messagges <- msg:
		fmt.Println("Sent messagges", msg)
	default:
		fmt.Println("no messagges sent")
	}

	select {
	case msg := <-messagges:
		fmt.Println("received messages", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
