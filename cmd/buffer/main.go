package main

import (
	"fmt"
	"time"
)

func sender(ch chan<- int, n int) {
	for i := 0; i <= n; i++ {
		ch <- i
		fmt.Println("sent:", i)
	}
	close(ch)
}

func receiver(ch <-chan int) {
	time.Sleep(1 * time.Second)
	for value := range ch {
		time.Sleep(1 * time.Second)
		fmt.Println("received:", value)
	}

	/*
		// alternative / uncommon way to receive values from channel
		for {
			value, ok := <-ch
			if !ok {
				break
			}
		}
	*/
}

func main() {
	fmt.Println("DEMO: unbuffered channel")
	unbufferedCh := make(chan int)
	go sender(unbufferedCh, 5)
	receiver(unbufferedCh)

	fmt.Println("DEMO: buffered channel")
	bufferedCh := make(chan int, 3)
	go sender(bufferedCh, 5)
	receiver(bufferedCh)
}
