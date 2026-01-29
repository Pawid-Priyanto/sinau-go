package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type cirlce struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c cirlce) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c cirlce) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g, "g")
	fmt.Println(g.area(), "g area")
	fmt.Println(g.perim(), "g perim")
	fmt.Println()
}

func detectCirlce(g geometry) {
	if c, ok := g.(cirlce); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func main() {
	fmt.Println("=========")
	r := rect{width: 3, height: 5}
	c := cirlce{radius: 14}

	measure(r)
	measure(c)

	detectCirlce(r)
	detectCirlce(c)
}
