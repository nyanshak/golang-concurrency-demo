package main

import "time"

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			ch <- i
		}
	}()

	time.Sleep(time.Second)
	close(ch)
	time.Sleep(time.Second)
}
