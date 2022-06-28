package main

import (
	"errors"
	"fmt"
)

func main() {
	foo()
}

func foo() error {
	defer func () {
		rec := recover()
		fmt.Printf("\nOh my god! We are in panic (%v)!\n\n", rec)
	} ()
	panic("you have been sanctioned")
	return errors.New("ohh, noo..")
}