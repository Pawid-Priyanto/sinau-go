package main

import (
	"fmt"
	"strconv"
)

func main() {

	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	g, _ := strconv.ParseInt("1234", 0, 64)
	fmt.Println(g)
	h, _ := strconv.ParseInt("0x1cb", 0, 64)
	fmt.Println(h)
	i, _ := strconv.ParseUint("1234", 0, 64)
	fmt.Println(i)
	j, _ := strconv.Atoi("1234")
	fmt.Println(j)

	_, e := strconv.Atoi("what")
	fmt.Println(e)
}
