package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("==== Slice ====")

	fmt.Println("unitunit: ", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp", s, "len", len(s), "cap:", cap(s))
	s = []string{"a", "b", "c"}
	fmt.Println("set", s)
	fmt.Println("get", s[2])

	fmt.Println("len", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println("s", s, "len:", len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c, "copy")

	l := s[2:5]
	fmt.Println("l", l)

	p := s[:5]
	fmt.Println(p, "ppp")
	r := s[2:]
	fmt.Println(r, "rrr")

	t := []string{"g", "h", "i"}
	t2 := []string{"g", "h", "i"}
	fmt.Println("ttt", t, t2)

	if slices.Equal(t, t2) {
		fmt.Println("t === t2", t, t2)
	}

	twoD := make([][]int, 3)

	fmt.Println(twoD)

	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d", twoD)

}
