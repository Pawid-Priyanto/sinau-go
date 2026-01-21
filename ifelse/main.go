package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even number")
	} else {
		fmt.Println("7 is odd number")
	}

	if 8%4 == 0 {
		fmt.Println("8 modulus 4 = 0")
	} else {
		fmt.Println("8 modulus 4 != 0")
	}
	if num := 9; num < 0 {
		fmt.Println(num, "is negatif")
	} else if num < 10 {
		fmt.Println(num, "is 1 digits")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// ============= FIZZBUZZ =============== //

	for n := 1; n <= 100; n++ {
		if n%3 == 0 && n%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if n%5 == 0 {
			fmt.Println("Buzz")
		} else if n%3 == 0 {
			fmt.Println("FIzz")
		} else {
			fmt.Println("loop", n)
		}
	}

}
