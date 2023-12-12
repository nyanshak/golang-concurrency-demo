package main

import (
	"sync"
	"time"
)

func DoWork(mu0 *sync.Mutex, mu1 *sync.Mutex) {
	mu0.Lock()
	defer mu0.Unlock()

	time.Sleep(2 * time.Second)

	mu1.Lock()
	defer mu1.Unlock()
}

func main() {
	mu0 := &sync.Mutex{}
	mu1 := &sync.Mutex{}

	go DoWork(mu0, mu1)
	DoWork(mu1, mu0)
}
