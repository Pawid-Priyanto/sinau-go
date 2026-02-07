package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("peach punch pecah"))
	fmt.Println("idx", r.FindStringIndex("peach punch pecah"))
	fmt.Println(r.FindStringSubmatch("peach punch pecah"))
	fmt.Println(r.FindStringSubmatchIndex("peach punch pecah "))
	fmt.Println(r.FindAllString("peach punch pinch poch", -1))
	fmt.Println("all: ", r.FindAllStringSubmatchIndex("peach punch pinch poch", -1))
	fmt.Println(r.FindAllString("peach punch pinch poch", 2))

	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruits>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
