package main

import (
	"fmt"
)

const (
	zero = iota // first value is zero
	one         // second value follows the sequence
	two         // so one
)

func main() {
	fmt.Printf("%v, %T", zero, zero)
	fmt.Printf("%v, %T", one, one)
	fmt.Printf("%v, %T", two, two)

}
