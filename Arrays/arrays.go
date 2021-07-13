package main

import (
	"fmt"
)

func main() {
	// Creates an array of lenght 3
	grades := [3]int{
		1, 2, 3,
	}
	fmt.Println(grades)

	// Instead use ... to shwo that the array is initialised with a fixed length
	num := [...]int{1, 2, 3}
	fmt.Println(num)

	// Declaration of Arrays and initialise seprately
	var students [3]string
	students[0] = "Pavan"
	students[1] = "Kumar"
	students[2] = "Gowda"
	fmt.Printf("Students:%v of lenght %v\n", students, len(students))

	//Matrix
	var matrix [3][3]int
	matrix[0] = [...]int{1, 2, 3}
	matrix[1] = [...]int{4, 5, 6}
	matrix[2] = [3]int{7, 8, 9}
	fmt.Println(matrix)

	// Passing arrays with refernce
	anotherMatrix := &matrix
	fmt.Println(*anotherMatrix)
	fmt.Println(&anotherMatrix) // Prints the memeory address

}
