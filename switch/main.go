package main

import (
	"fmt"
	"time"
)

func main() {

	n := 2

	switch n {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its weekend")
	default:
		fmt.Println("Its Weekday")
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Its Before Noon")
	default:
		fmt.Println("Its After Noon")
	}

	whatAmI := func(oi interface{}) {
		switch t := oi.(type) {
		case bool:
			fmt.Println(" I am Boolean")
		case int:
			fmt.Println(" I am Integer")
		case string:
			fmt.Println(" I am String", t)
		}
	}
	whatAmI(true)
	whatAmI(23)
	whatAmI("logic")

}
