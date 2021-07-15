package main

import (
	"fmt"
	"log"
	"net/http"
)

//******* Note :  To understand this please run the program  twice so that the port is not available and program panics *************
func main() {
	// Defer always execute at the end and in LIFO order
	fmt.Println("First")
	defer fmt.Println("Last")
	fmt.Println("Second")

	defer fmt.Println("Fifth")
	defer fmt.Println("Fourth")

	// Run the server twice to see the program panic
	Server()
	fmt.Println("Third")

	// Panic occur after the defer statements are executed

}

func Server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		// To recover from a panic use the below way , to check how the program panic's comment out the next 4 lines
		defer func() {
			if err := recover(); err != nil {
				log.Println("Recovering but no second server")
			}
		}()
		panic("Panic here due to port being used")
	}
}
