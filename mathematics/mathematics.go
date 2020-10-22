package main

import "fmt"

func main() {
	a := 8
	fmt.Println(a << 3) // multiplied by 2 3 times (2ˆ3 * 2ˆ3 * 2ˆ3)
	fmt.Println(a << 4) // multiplied by 2 4 times (2ˆ3 * 2ˆ3 * 2ˆ3 * 2ˆ3)
	fmt.Println(a << 5) // multiplied by 2 5 times (2ˆ3 * 2ˆ3 * 2ˆ3 * 2ˆ3 * 2ˆ3)

	fmt.Println(a >> 3) // Divided by 2 3 times (2ˆ3 / 2ˆ3 / 2ˆ3)
	fmt.Println(a >> 4) // Divided by 2 4 times (2ˆ3 / 2ˆ3 / 2ˆ3 / 2ˆ3)
	fmt.Println(a >> 5) // Divided by 2 5 times (2ˆ3 / 2ˆ3 / 2ˆ3 / 2ˆ3 / 2ˆ3)

}
