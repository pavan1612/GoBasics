package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var args []string = os.Args
	// To convert string argument to int
	i, _ := strconv.Atoi(args[1])

	// Break keyword is not needed to be specified , it is inbuilt .
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Another number")
	}
	// Can use break to exit inbetween a case block . Note : case keyword acts as a limiters between case blocks
	// Switch with initialization and condition . Can Use multiple case condition in a single case block
	switch j := i + 5; j {
	case 6, 7, 8:
		fmt.Println("Adding five gives six , seven or eight ")
	case 9, 10, 11:
		fmt.Println("Adding five gives nine , ten or eleven")
	default:
		fmt.Println("Not in range ")
	}

	// Switch without any condition or initilisation
	switch {
	case i < 2:
		fmt.Println("one less than two ")
	case i < 3: // Note that even if the below cases also is valid it doesnt execute , To overcome use fallthrough
		fmt.Println("two less than five ")
	default:
		fmt.Println("Greater than three  ")
	}

	// Can use break to exit inbetween a case block . Note : case keyword acts as a limiters between case blocks

}
