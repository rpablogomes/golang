package main

import "fmt"

func main() {
	a := 10 // 1010
	b := 2  //0010

	fmt.Println(a & b)  // 0010
	fmt.Println(a | b)  // 1010
	fmt.Println(a ^ b)  // 1000
	fmt.Println(a &^ b) // 1000
}
