package main

import "fmt"

func main() {
	fmt.Println("===== Array =====")

	var a [5]int
	fmt.Println("array ", a)

	a[4] = 100
	fmt.Println("set array ke4", a)
	fmt.Println("get array ke4", a[4])
	fmt.Println("length", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("array b", b)

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("arry b with...", b)

	b = [...]int{100, 3: 400, 500}
	fmt.Println("array with indexing", b) // [100, 0, 0, 400, 500]

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = j + i
		}
	}

	// i = 0 0 0 1 1 1
	// j =  0 1 2 0 1 2
	// res = [0, 1, 2], [1, 2, 3]

	fmt.Println("twoD", twoD)

	var twoDD = [2][3]int{
		{1, 2, 3}, {1, 2, 3},
	}
	fmt.Println(twoDD, "assign array twod")

}
