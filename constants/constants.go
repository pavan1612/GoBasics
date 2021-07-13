package main

import (
	"fmt"
	"math"
)

// Declaring a string with global
const s string = "Pavan is constant"

func main() {

	fmt.Println(s)

	// Constants can be used where variables are used - Here printing n
	const n = 5000000000
	fmt.Printf("%T\n", n)

	const d = 3e10 / n
	fmt.Println(int(d))
	fmt.Println(math.Sin(n))

	// Enums in GO is done through constant block using "iota"
	// iota starts from 0 and scoped to the constant block
	const (
		a = iota
		b
		c
	)
	fmt.Println(a, b, c)

	// Here we declare the one element as _(Write only variable) and understant enums
	// We can also add an offset or use bitwise to raise power to the power of 2
	const (
		_ = iota + 5
		can
		dog
		elephant
	)
	var animal int = dog
	fmt.Printf("%v\n", animal == dog)
	fmt.Println(dog)

	// Here the isAdmin is 0001 and isHeadQuarters is 0010 amd so on due to bit shifting operation <<
	const (
		isAdmin = 1 << iota
		isHeadQuarters
		canSeeFinancials
		canSeeProjectA
	)
	var CEORole = isAdmin | isHeadQuarters | canSeeProjectA
	fmt.Printf("%b\n", CEORole) //  Output for this would be 1101 since CEO cant see financials

	// While using if/else there is two types one is initializer/condition ; result/consition or conditon
	// Very Important to remember that else comes in the same line as if's block close
	if result := isAdmin&CEORole == isAdmin; result {
		fmt.Println("Has Admin Access")
	} else {
		fmt.Println("Doesnt have admin access")
	}
}
