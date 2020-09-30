package main

import "fmt"

func main() {
	first := 1 == 1
	second := 1 == 2

	fmt.Printf("%v, %T", first, first)
	fmt.Printf("%v, %T", second, second)

}
