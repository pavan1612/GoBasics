package main

import (
	"fmt"
	"reflect"
)

// Creating a struct for Employee similar to creating a class
type Employee struct { // Capital letter for Employee meaning it can be exported
	id       int `required:"Yes"` // Can have Tags like this but it doesnt mean anything to the compiler other than text
	name     string
	isActive bool
	team     []string
}

func main() {

	// Creating an Employee similar to creating an object in OOP
	pavan := Employee{
		id:       1, // Consider having field name(id) while assigning value
		name:     "pavankumar",
		isActive: true,
		team: []string{
			"Goutham",
			"Eric",
			"JK",
			"Peter",
		},
	}

	// Second way of initialising an object
	pavan2 := Employee{}
	pavan2.id = 2
	pavan2.name = "pavankumar"
	pavan2.isActive = true

	// If an empty object is created then be careful as values are set to 0
	kumar := Employee{}
	fmt.Println(pavan)
	fmt.Printf("Name is %v and active status is %v with teammates %v\n", pavan.name, pavan.isActive, pavan.team)
	fmt.Printf("Kumar Id is %v\n", kumar.id) // will return 0 when empty

	// To check elements are present
	if kumar.id == 0 || kumar.name == "" {
		fmt.Println("Element absent")
	}

	// Creating an ananymous struct Note : can be used as a struct for immeddiate incoming data to pass on
	pavana := struct{ name string }{
		name: "pavan",
	}
	fmt.Println(pavana)

	// Objects created are value type and are copied and not referneced like structs and maps
	p := pavana
	p.name = "kumar"
	fmt.Println(pavana, p)

	// Use reflect package to know the properties of a field in a struct like Tag here
	t := reflect.TypeOf(Employee{})
	field, _ := t.FieldByName("id")
	fmt.Println(field.Tag)
	fmt.Println(t)

}
