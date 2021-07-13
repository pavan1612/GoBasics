package main

import (
	"fmt"
)

func main() {
	// Slices are reference type of arrays
	a := []int{1, 2, 3, 4}
	b := a
	b[1] = 10
	fmt.Println(a, b)
	fmt.Printf("Length : %v\nCapacity : %v\n", len(a), cap(a))

	// Createing slices form an arrays Note : not literally
	aa := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	bb := aa[:]
	cc := aa[3:]  // First number is inclusive
	dd := aa[3:8] // Second number is exclusive
	fmt.Printf("All the Slices: %v\t%v\t%v\t%v\n", aa, bb, cc, dd)

	// Using make to create slices
	aaa := make([]int, 2)
	aaa[0] = 1
	aaa[1] = 2
	fmt.Printf("Using make : %v\n", aaa)
	fmt.Printf("Value: %v\n", len(aaa))
	fmt.Printf("Capacity: %v\n", cap(aaa))

	// To understand how append is used and how it impacts the length and capacity of the slice
	aaa = append(aaa, 3, 4, 5)
	fmt.Printf("Value: %v\n", len(aaa))
	fmt.Printf("Capacity: %v\n", cap(aaa))

}
