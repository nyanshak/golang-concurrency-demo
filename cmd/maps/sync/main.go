package main

import (
	"fmt"
	"sync"

	"github.com/nyanshak/golang-concurrency-demo/pkg/maps"
)

func main() {
	myMap := map[string]string{}
	mutex := &sync.Mutex{}

	var wg0 sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg0.Add(1)
		go func() {
			maps.WriteMapMutex(myMap, "some_key", "some_value", mutex)
			wg0.Done()
		}()
	}
	wg0.Wait()
	fmt.Println("finished writing to map (with mutex) successfully")

	var safeMap sync.Map
	var wg1 sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg1.Add(1)

		go func() {
			maps.WriteSyncMap(&safeMap, "some_key", "some_value")
			wg1.Done()
		}()
	}
	wg1.Wait()
	fmt.Println("finished writing to sync.Map successfully")
}
