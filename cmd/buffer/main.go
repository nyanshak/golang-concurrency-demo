package main

import (
	"fmt"
	"time"
)

func sender(ch chan int) {
	for i := 0; i <= 9; i++ {
		ch <- i
		fmt.Println("sent:", i)
	}
	close(ch)
}

func main() {
	unbufferedCh := make(chan int)
	bufferedCh := make(chan int, 3)

	go sender(unbufferedCh)
	time.Sleep(1 * time.Second)
	fmt.Println("unbuffered channel values:")
	for value := range unbufferedCh {
		fmt.Println("received:", value)
	}

	go sender(bufferedCh)
	time.Sleep(1 * time.Second)
	fmt.Println("buffered channel values:")
	for value := range unbufferedCh {
		fmt.Println("received:", value)
	}
}
