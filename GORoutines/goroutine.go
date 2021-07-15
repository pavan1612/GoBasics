package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	// runtime operation to set the number of threads to operate on
	runtime.GOMAXPROCS(100)
	// A  simple go routine
	go printHello()
	time.Sleep(10 * time.Millisecond) // Pausing the program to execute the thread - Bad way of synchronizing

	// A go routine to indicate the right way to pass values into a dunction executed as a go routine
	s := "Pavan"
	go func(s string) {
		fmt.Println(s)
		printHello()
	}(s)
	time.Sleep(100 * time.Millisecond) // Pausing the program to execute the thread - Bad way of synchronizing

	// Synchronised way to execute a go routine using waitgroup and read write mutex
	for i := 0; i < 10; i++ {
		wg.Add(2) // Adds 2 count into the stack of waitgroup
		m.RLock() // Locks the execution of any reading operation in the main program so that below routine can execute
		go printCounter()
		m.Lock() // Locks the execution of any writing operation in the main program
		go increment()
	}
	wg.Wait() // Acts like pausing the program(like time.sleep) until the waitgroup stack is empty
}
func printHello() {
	fmt.Println("Hello")

}

func printCounter() {
	fmt.Println(counter)
	m.RUnlock() // Unlocks read operation
	wg.Done()   // deletes one count from the waitgroup stack
}
func increment() {
	counter++
	m.Unlock() // Unlocks write operation
	wg.Done()  // deletes one count from the waitgroup stack
}
