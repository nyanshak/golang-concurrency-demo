package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan<- int, n int, shouldCausePanic bool) {
	for i := 0; i < n; i++ {
		fmt.Println("waiting to send:", i)
		ch <- i
		fmt.Println("sent:", i)
		time.Sleep(1 * time.Second)
	}

	// If multiple producers are writing to the channel, the channel should be
	// closed only once. If the channel is closed multiple times, it will
	// panic.
	if shouldCausePanic {
		fmt.Println("sender: closing channel")
		close(ch)
	}
}

func consumer(ch <-chan int) {
	for value := range ch {
		fmt.Println("received:", value)
	}
	fmt.Println("receiver: channel closed")
}

func demo(numProducers int, jobsPerProducer int, shouldCausePanic bool) {
	ch := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < numProducers; i++ {
		wg.Add(1)
		go func() {
			producer(ch, jobsPerProducer, shouldCausePanic)
			wg.Done()
		}()
	}

	go consumer(ch)

	wg.Wait()
	close(ch)
}

func main() {
	producerCount := 2
	jobsPerProducer := 3
	shouldCausePanic := false
	demo(producerCount, jobsPerProducer, shouldCausePanic)

	// Wait for the goroutines to finish.
	time.Sleep(3 * time.Second)

	fmt.Printf("\nDEMO: cause panic by closing sender channel multiple times")
	jobsPerProducer = 1
	shouldCausePanic = true
	demo(producerCount, jobsPerProducer, shouldCausePanic)
}
