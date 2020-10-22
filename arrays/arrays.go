package main

import "fmt"

func main() {
	var students = [3]int{1, 2, 3} //[3] is the amout of elements in the array. int is the data type. Numbers in {} are he elements of the array.
	fmt.Printf("Students: %v", students)

	var teachers = [3]string{} // The array is empty
	teachers[0] = "John"       //Declaration of value on first positon of the array teachers
	teachers[1] = "Paul"       //Declaration of value on second positon of the array teachers
	teachers[2] = "Wilson"     //Declaration of value on third positon of the array teachers

	fmt.Printf("teachers: %v", teachers)

}
