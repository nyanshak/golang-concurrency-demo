package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func incrementNoSynchronization(times int) {
	n := 0
	var wg sync.WaitGroup
	for i := 0; i < times; i++ {
		wg.Add(1)
		go func() {
			n++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("incremented %d times - final value: %d\n", times, n)
}

func incrementWithMutex(times int) {
	n := 0
	var wg sync.WaitGroup
	m := sync.Mutex{}
	for i := 0; i < times; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m.Lock()
			defer m.Unlock()
			n++
		}()
	}
	wg.Wait()
	fmt.Printf("incremented %d times - final value: %d\n", times, n)
}

func incrementWithAtomic(times int) {
	var n int64
	var wg sync.WaitGroup
	for i := 0; i < times; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&n, 1)
		}()
	}
	wg.Wait()
	fmt.Printf("incremented %d times - final value: %d\n", times, n)
}

func main() {
	fmt.Println("\n~~~~~ NO SYNCHRONIZATION ~~~~~")
	for i := 0; i < 5; i++ {
		incrementNoSynchronization(1000)
	}

	fmt.Println("\n~~~~~ WITH MUTEX ~~~~~")
	for i := 0; i < 5; i++ {
		incrementWithMutex(1000)
	}

	fmt.Println("\n~~~~~ WITH ATOMIC ~~~~~")
	for i := 0; i < 5; i++ {
		incrementWithAtomic(1000)
	}
}
