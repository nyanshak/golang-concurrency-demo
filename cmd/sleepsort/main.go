package main

import (
	"fmt"
	"sync"
	"time"
)

func sleepAndPrint(n int64, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(n) * time.Millisecond)
	fmt.Println("\t", n)
}

func Sort(numbers []int64) {
	var wg sync.WaitGroup

	wg.Add(len(numbers))

	for _, n := range numbers {
		go sleepAndPrint(n, &wg)
	}

	wg.Wait()
}

func demonstrate(name string, numbers []int64) {
	fmt.Printf("###### %s #####\n", name)
	fmt.Println("Unsorted")
	for _, n := range numbers {
		fmt.Println("\t", n)
	}
	fmt.Println("Sorted")
	Sort(numbers)
}

func main() {
	case0 := []int64{5, 3, 6, 3, 1, 2, 7, 9, 8, 4}
	case1 := []int64{5, 3, 6, 3, 1, 2, 7, 9, 8, 4}
	case2 := []int64{5, 3, 6, 3, 1, 2, 7, 9, 8, 4, 123, 17, 456, 55, 7, 19, 1, 2}

	demonstrate("Case 0", case0)
	demonstrate("Case 1", case1)
	demonstrate("Case 2", case2)
}
