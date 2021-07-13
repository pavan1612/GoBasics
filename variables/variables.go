package main

import "fmt"

func main() {

	// Any variable without a value assigned is - 0
	var z int64
	fmt.Println(z)

	//Declares and  prints the integer with full decalration
	var i, j int8 = 1, 2
	fmt.Println(i, j)

	// Arithmatic Operations like (+,-,*,/,%) can be done
	fmt.Println(i % j)

	// can also have 16/variable bits intergers
	var i16 int16 = 10000
	fmt.Println(i16)

	//Declares and  prints the 32 bits float
	var f32 float32 = 1.1111
	fmt.Println(f32)

	//Declares and  prints a 64 bit complex number with 32 bits for real part and 32 bits for imaginary part
	var comp complex64 = 1 + 55i
	fmt.Println(comp)

	// another way of declaration for variables
	num := 44.5
	fmt.Println(num)

	//Declares and prints a string
	str := "Pavan"
	fmt.Println(str)
	fmt.Println(len(str))

	//Declares and pritns a string in real form while ignoring any inbuilt identifiers
	str2 := `
	Hey Hello 
		everyone 
			See the new lines and tab spacing 
	
			here i am 
	`
	fmt.Println(str2)

	// Declaring and printing the boolean value
	status := true
	fmt.Println(status)

	//Declares and prints a rune - rune is a ascii value equvalent of a character
	var r rune = 'a'
	fmt.Println(r)
	// To print the String equivalent of int32
	fmt.Println(string(r))

}
