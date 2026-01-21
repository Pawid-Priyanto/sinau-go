package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {

	fmt.Println("constant s", s)
	const n = 500000

	fmt.Println(n)

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
