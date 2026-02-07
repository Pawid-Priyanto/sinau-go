package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}
func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		panic(err)
	}
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func myPanic() {
	panic("there is a problem")
}

func main() {
	// panic("a problem")

	// path := filepath.Join(os.TempDir(), "file")
	// _, err := os.Create(path)
	// if err != nil {
	// 	panic(err)
	// }

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recoverd error:\n", r)
		}
	}()

	myPanic()

	fmt.Println("After myPanic()")
	path := filepath.Join(os.TempDir(), "defer.txt")
	f := createFile(path)
	defer closeFile(f)
	writeFile(f)
}
