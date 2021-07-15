package main

import (
	"fmt"
)

func main() {

	// Simple for loop
	for i := 0; i < 5; i++ {
		fmt.Printf("\ni: %v", i)
	}
	// For loop like while loop (Different Usage ) also showing the use of continue keyword
	j := 0
	for j < 10 {
		j++
		if j%2 == 0 {
			continue
		}
		fmt.Printf("\nj : %v", j)
	}
	// For loop with no condition and showing the usage of break
	k := 0
	for {
		fmt.Printf("\nk : %v", k)
		k++
		if k > 5 {
			break
		}
	}

	// Multiple for loop showing the use of goto function using a break labels
Loop:
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			fmt.Printf("\ni*j : %v ", i*j)
			if i*j >= 5 {
				break Loop
			}
		}
	}

	// For loop used as a for each by looping through collection using range
	s := []int{1, 2, 3}
	fmt.Println()
	for key, value := range s {
		fmt.Println(key, value)
	}

}
