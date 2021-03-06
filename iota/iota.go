package main

import (
	"fmt"
)

const (
	zero = iota // first value is zero
	one         // second value follows the sequence
	two         // so on
)

const (
	five = iota + 5 // Value starts from 5
	six             //so on
)

func main() {
	fmt.Printf("%v, %T", zero, zero)
	fmt.Printf("%v, %T", one, one)
	fmt.Printf("%v, %T", two, two)

	fmt.Printf("%v, %T", five, five)
	fmt.Printf("%v, %T", six, six)
}
