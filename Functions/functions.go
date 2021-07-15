package main

import (
	"fmt"
)

func main() {
	// Returns two values
	fmt.Println(twoReturnValues(true))
	// Can pass multiple parameters
	multipleParameters("Values are", 12, 13, 14, 15)
	// Anonynous functions
	multiply := func(a int, b int) int {
		return a * b
	}
	fmt.Printf("Multiplied answer is : %v", multiply(44, 55))
	// To retunr a named return value
	fmt.Println("Named Retrun Value ", namedReturn())
	// To pass a value by reference
	a := 10
	fmt.Printf("Value of 'a' before function call : %v\n", a)
	referenceFunc(&a)
	fmt.Printf("Value of 'a' after referencing : %v\n", a)
	// To pass functions around 
	passFunction(multiply)

}

func twoReturnValues(condition bool) (string, string) {
	if condition {
		return "Error", "Error body"
	}
	return "Result", "No Error"
}

func multipleParameters(s string, values ...int) {
	fmt.Println(s)
	for _, v := range values {
		fmt.Println(v)
	}
}
func namedReturn() (result string) {
	result = "Result value "
	return
}
func referenceFunc(b *int) {
	*b = 20
}
func passFunction(mul func(int, int) int) {
	fmt.Println(mul(10, 30))
}
