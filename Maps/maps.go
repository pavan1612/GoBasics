package main

import (
	"fmt"
)

func main() {
	// Shows how to create a simple map
	// Slices cant be a key type
	statePopulation := map[string]int{
		"VIC": 10000,
		"NSW": 20000,
		"NT":  5000,
		"TAS": 7000,
		"WA":  8000,
	}
	fmt.Printf("Population of all states is %v\n", statePopulation)
	fmt.Printf("Victoria: %v\n", statePopulation["VIC"]) // To access individual element

	// Using make to declare a map to be used later
	m := make(map[int]int)
	m = map[int]int{
		1: 8,
		2: 16,
		3: 24,
	}
	m[4] = 32 // Note when adding an element to map ordering is not guaranteed
	fmt.Printf("Number of bits in 2 byte is %v\n", m[2])
	fmt.Println(m)

	// To delete an element form a map
	delete(m, 2)
	fmt.Printf("After deletion : %v\n", m)

	// Maps for any not present in the map returns 0 , to avoid error we can check if a key present
	_, ok := m[2]
	fmt.Println(ok) // retuns false because of deletion

	// Similar to Slices any chnages while copied affects all the copies because only references are passed
	sp := statePopulation
	delete(sp, "TAS")
	fmt.Println(sp)
	fmt.Println(statePopulation)
	

}
