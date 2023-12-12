package main

import (
	"fmt"
	"time"
)

// TODO: direction of channel
func producer(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("sent:", i)
		time.Sleep(1 * time.Second)
	}
	// TODO: close channel
}

func consumer(ch chan int) {
	// TODO: range over channel
	for {
		value, ok := <-ch
		if !ok {
			fmt.Println("channel closed")
			break
		}
		fmt.Println("received:", value)
	}
}

func main() {
	ch := make(chan int)
	go consumer(ch)
	go producer(ch)

	// Sleep to let goroutines finish
	time.Sleep(7 * time.Second)
}
