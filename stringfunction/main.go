package main

import (
	"fmt"
	"os"
	s "strings"
)

var p = fmt.Println

type point struct {
	x, y int
}

func main() {
	p("Contains: ", s.Contains("test", "es"))
	p("Count: ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSufix: ", s.HasSuffix("test", "st"))
	p("Index: ", s.Index("test", "e"))
	p("Join: ", s.Join([]string{"test", "es"}, "-"))
	p("Replace: ", s.Replace("test", "es", "tes", -1))
	p("Repeat: ", s.Repeat("test", 5))
	p("Split: ", s.Split("t-e-s-t", "-"))
	p("ToLower: ", s.ToLower("TeSt"))
	p("ToUpper: ", s.ToUpper("test"))

	// ===== string formating ===== //
	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)
	fmt.Printf("struct2: %+v\n", p)
	fmt.Printf("struct3: %#v\n", p)

	fmt.Printf("type: %T\n", p)
	fmt.Printf("bool: %t\n", true)
	fmt.Printf("int: %d\n", 123)
	fmt.Printf("bin: %b\n", 14)
	fmt.Printf("char: %c\n", 33)
	fmt.Printf("hex: %x\n", 456)
	fmt.Printf("float1: %f\n", 78.9)
	fmt.Printf("float2: %e\n", 1233394.0)
	fmt.Printf("float3: %E\n", 12343423900.0)

	fmt.Printf("str1: %s\n", "\"string\"")
	fmt.Printf("str1: %q\n", "\"string\"")
	fmt.Printf("str1: %x\n", "hex this")
	fmt.Printf("pointer: %p\n", &p)

	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "boo")
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "boo")

	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")

}
