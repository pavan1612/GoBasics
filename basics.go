package main

import "fmt"


// Defining a struct of type "person" and assigning the properties and methods  
type person struct {
	name string
	age  int
}
func (p person) speak() {
	fmt.Println("This is  " + p.name + " speaking ")
}

// Defining a struct of type "hitman" and assigning the properties and methods  
type hitman struct {
	p       person
	canKill bool
}
func (h hitman) speak() {
	fmt.Println("Hitman here: " + h.p.name)
}


// Defining a interface of type "human" and assigning the properties and methods  
type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}



// "main" function to understand the working  
func main() {

	s1 := []int{1, 2, 3}
	fmt.Println(s1)

	m1 := map[string]int{
		"Pavan":   25,
		"Kishore": 21,
	}
	fmt.Println(m1)

	pavan := person{
		"Pavan ",
		25,
	}

	kishore := person{
		"Kishore",
		21,
	}
	hit := hitman{
		pavan,
		true,
	}
	// Using methods
	pavan.speak()
	kishore.speak()
	hit.speak()
	// Using Interface
	saySomething(pavan)
	saySomething(hit)

}
