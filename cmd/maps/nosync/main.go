package main

import (
	"fmt"
	"sync"

	"github.com/nyanshak/golang-concurrency-demo/pkg/maps"
)

func main() {
	myMap := map[string]string{}

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			maps.WriteMapNoSync(myMap, "some_key", "some_value")
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("got lucky - successfully wrote to map concurrently without synchronization")
}
