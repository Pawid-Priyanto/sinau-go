package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("string and rune")

	const s = "ꦥꦮꦶꦠ꧀ ꦥ꧀ꦫꦶꦪꦤ꧀ꦠꦺꦴ"
	// const s = "สวัสดี"
	// const s = "hello"

	fmt.Println("len:", len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println()
	fmt.Println("Rune count: ", utf8.RuneCountInString(s))

	for idx, runValue := range s {
		fmt.Printf("%#U starts at %d\n", runValue, idx)
	}
	fmt.Println("\nUsing Decode Rune In String")
	for i, w := 0, 0; i < len(s); i += w {
		runValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runValue, i)
		w = width

		examineRuneValue(runValue)
	}
}

func examineRuneValue(r rune) {
	switch r {
	case 'h':
		fmt.Println("Found hello")
	case 'ꦥ':
		fmt.Println("found aksara")
	}
}
