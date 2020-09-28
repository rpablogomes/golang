package main

import "fmt"

var first, second, third = 1, 2, 3 // Way to declare multiple variables. GLOBAL
var fourth int = 4                 // Common way to declare a varible with data type(int). GLOBAL

var test = false //Boolean. GLOBAL

func main() {
	fifth := 5 // Short way to declare a variable. ONLY FOR LOCAL VARIABLES

	fmt.Println(first, second, third)
	fmt.Printf("%v, %T", fourth, fourth) //("VALUE, TYPE", value of fourth, type of forth)
	fmt.Printf("%v, %T", fifth, fifth)   //("VALUE, TYPE", value of fourth, type of forth)
	fmt.Printf("%v, %T", false, false)   //("VALUE, TYPE", value of fourth, type of forth)
}
