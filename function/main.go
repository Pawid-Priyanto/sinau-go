package main

import "fmt"

func main() {
	fmt.Println("====== fuction ======")
	res := add(34, 23)
	res1 := plusPlus(1, 2, 3)
	fmt.Println(res, "res1", res1)

	// =========== multiple function ============ //
	a, b := multiple()
	fmt.Println(a, b)
	_, c := multiple()
	fmt.Println(c)
	// =========== Variadic Function ============= //
	sum(1, 2)
	sum(1, 2, 3)

	num := []int{1, 2, 3, 4, 5}
	sum(num...)

	// ========== closures =========== //
	nextInt := intSeq()
	fmt.Println(nextInt(), "intseq ====")
	fmt.Println(nextInt(), "intseq ====")
	fmt.Println(nextInt(), "intseq ====")

	newsIntSeq := intSeq()
	fmt.Println(newsIntSeq(), "new sequent")

	// ========== recrusive ============ //
	fact1 := fact(0)
	fmt.Println(fact1, "fact 1")
	fact7 := fact(7)
	fmt.Println(fact7, "fact7")
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))

	// ========= range over built in types ========= //
	// rangeOver := rangeover()

	kalkulasi()
}

func add(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func multiple() (int, int) {
	return 3, 7
}

func sum(nums ...int) {
	fmt.Print(nums, " ")

	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// recrusive
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

// range over built in types
func kalkulasi() {
	numb := []int{2, 3, 4}
	sum := 0

	for _, num := range numb {
		sum += num
	}

	fmt.Println(sum, "sum")

	for i, num := range numb {
		if num == 3 {
			fmt.Println("index", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
