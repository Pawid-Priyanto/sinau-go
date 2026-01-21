package main

import "fmt"

func main() {
	fmt.Println("===================")

	i := 1
	for i <= 3 {
		fmt.Println("Print i", i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println("Print j =", j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for k := range 3 {
		fmt.Println("Print K", k)
	}
	for l := range 6 {
		if l%2 == 0 {
			continue
		}
		fmt.Println("Print l", l)
	}
	fmt.Println("===================")

}
