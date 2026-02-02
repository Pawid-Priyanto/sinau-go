package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("cannot work with 42")
	}

	return arg + 3, nil
}

var ErrorOutOfTea = errors.New("No More Tea avaliable")
var ErrorPower = errors.New("cannot boil water")

func makeTea(arg int) error {
	switch arg {
	case 2:
		return ErrorOutOfTea
	case 4:
		return fmt.Errorf("making tea: %w", ErrorPower)
	}
	return nil
}

type argErr struct {
	arg     int
	message string
}

func (e *argErr) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func l(arg int) (int, error) {
	if arg == 42 {
		return -1, &argErr{arg, "Can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f(i); e != nil {
			fmt.Println("f failed", e)
		} else {
			fmt.Println("f work", r)

		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrorOutOfTea) {
				fmt.Println("We should buy a new Tea")
			} else if errors.Is(err, ErrorPower) {
				fmt.Println("Now it is dark")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}
		fmt.Println("Tea is Ready")
	}

	_, err := l(41)
	var e *argErr
	if errors.As(err, &e) {
		fmt.Println(e.arg)
		fmt.Println(e.message)
	} else {
		fmt.Println("error doesnot match arg error")
	}
}
