package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomErrorWorker(id int, resultCh chan<- int, errorCh chan<- error) {
	for {
		time.Sleep(time.Second)

		// pick a random number between 1 and 10
		x := rand.Intn(10) + 1
		if x%5 == 0 {
			errorCh <- fmt.Errorf("something went wrong in worker %d", id)
			continue
		}

		resultCh <- x
	}
}

func main() {
	resultCh := make(chan int)
	errorCh := make(chan error)

	for i := 0; i < 5; i++ {
		go randomErrorWorker(i, resultCh, errorCh)
	}

	for {
		select {
		case result := <-resultCh:
			fmt.Println("result:", result)
		case err := <-errorCh:
			fmt.Println("error:", err)
		}
	}
}
