package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func consumeFromQueue(queue <-chan int, wg *sync.WaitGroup) {
	for i := range queue {
		// sleep for a random amount of time between 1 and 3 seconds
		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)

		wg.Done()
		fmt.Println("Processed", i)
	}
}

func main() {
	numWorkers := 5

	numJobs := 500
	var wg sync.WaitGroup
	wg.Add(numJobs)

	queue := make(chan int, numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go consumeFromQueue(queue, &wg)
	}

	for j := 0; j < numJobs; j++ {
		queue <- j
	}

	wg.Wait()
	fmt.Println("All jobs done")
}
