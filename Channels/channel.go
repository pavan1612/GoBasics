package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	// Creating two channels
	ch := make(chan int, 100)
	chStr := make(chan string)
	// Creating go routines and writing channels
	start := time.Now()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			ch <- i
			wg.Done()
		}(i)
	}

	// Creating go routines and receiving channels
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()
	}

	// To make the go routine receive-only/send-only channel specific
	wg.Add(1)
	go func(chString chan<- string) {
		chString <- "Pavan"
		wg.Done()
	}(chStr)

	wg.Add(1)
	go func(chString <-chan string) {
		str := <-chStr
		fmt.Println(str)
		wg.Done()
	}(chStr)
	wg.Wait()
	GoTime := time.Since(start)

	// Sequential execution of the same program
	start = time.Now()
	arr := [100]int{}
	for i := 0; i < 100; i++ {
		arr[i] = i
	}
	for _, v := range arr {
		fmt.Println(v)
	}
	SequentialTime := time.Since(start)

	// In this program both go routines and Sequential are doing copying operations , so no performance improvement is observed
	// Goroutine makes this program slower due to spwanning new go routines
	log.Printf("Go routines with Channels took :%v", GoTime)
	log.Printf("Time elapsed for sequential copying of data and printing : %v", SequentialTime)

}
