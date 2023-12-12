package main

import (
	"fmt"
	"time"
)

func sender(ch chan<- string, waitDuration time.Duration, senderName string) {
	for {
		ch <- fmt.Sprintf("message from %s", senderName)
		time.Sleep(waitDuration)
	}
}

func receiver(ch1 <-chan string, ch2 <-chan string) {
	ticker := time.NewTicker(1250 * time.Millisecond)

	for {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		case <-ticker.C:
			fmt.Println("ticker interval")
		}
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go sender(ch1, 2*time.Second, "ONE")
	go sender(ch1, 3*time.Second, "TWO")
	receiver(ch1, ch2)
}
