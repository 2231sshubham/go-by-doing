package main

import (
	"fmt"
)

func printEven(odd, even chan int, done chan struct{}) {
	// defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		<-even
		fmt.Print(i, " ")
		if i == 10 {
			done <- struct{}{}
			return
		}
		odd <- 1
	}
}

func printOdd(odd, even chan int, done chan struct{}) {
	// defer wg.Done()
	for i := 1; i < 10; i += 2 {
		<-odd
		fmt.Print(i, " ")
		even <- 1
	}
}

func main() {
	// var wg sync.WaitGroup
	// wg.Add(2)

	done := make(chan struct{})
	oddChan := make(chan int)
	eveChan := make(chan int)

	go printEven(oddChan, eveChan, done)
	go printOdd(oddChan, eveChan, done)

	oddChan <- 1
	// wg.Wait()
	<-done
}
