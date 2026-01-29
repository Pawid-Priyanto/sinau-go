package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func newPerson(name string) *Person {
	p := Person{name: name}
	p.age = 42
	return &p
}

func main() {

	fmt.Println(Person{"harrop", 23})
	fmt.Println(Person{name: "ahmad", age: 27})
	fmt.Println(Person{name: "liam"})
	fmt.Println(&Person{name: "Dean", age: 31})
	fmt.Println(newPerson("Jhon"))

	s := Person{name: "arne", age: 15}
	fmt.Println(s.age, s.name)

	sp := &s
	fmt.Println("sp:", sp.age)

	dog := struct {
		name   string
		isBool bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)

}
