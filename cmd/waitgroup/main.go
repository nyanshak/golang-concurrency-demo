package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("worker %d starting\n", id)

	// Simulate doing some expensive task by sleeping a random time betwen 1-5
	// seconds.
	time.Sleep(time.Second * time.Duration(rand.Intn(5)+1))

	fmt.Printf("worker %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("all workers completed their work")
}
