package main

import (
	"fmt"
	"maps"
)

func main() {
	list := make(map[string]int)
	fmt.Println(list, "list")

	list["k1"] = 7
	list["k2"] = 14

	fmt.Println(list, "2 list")

	v1 := list["k2"]
	fmt.Println("v1", v1)

	v2 := list["k3"]
	fmt.Println("v2", v2)

	fmt.Println("len:", len(list))

	delete(list, "k2")
	fmt.Println("list", list)
	clear(list)
	fmt.Println("delete list", list)

	_, prs := list["k2"]
	fmt.Println(prs, "prs")
	n := map[string]int{"foo": 1, "bar": 2}
	n2 := map[string]int{"foo": 1, "bar": 2}

	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
