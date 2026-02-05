package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	// sorting
	strings := []string{"c", "a", "b"}
	slices.Sort(strings)
	fmt.Println("string:", strings)

	number := []int{2, 9, 6}
	slices.Sort(number)
	fmt.Println("number", number)

	isSorted := slices.IsSorted(number)
	fmt.Println(isSorted)

	// sorting by function

	fruits := []string{"peach", "kiwi", "banana"}

	lenCamp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCamp)

	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		{name: "heru", age: 47},
		{name: "budi", age: 17},
		{name: "alex", age: 34},
	}

	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})

	fmt.Println(people)

}
