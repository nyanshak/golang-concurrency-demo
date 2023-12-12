package main

import (
	"fmt"
	"time"
)

func main() {
	c0 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		c0 <- "result 1"
	}()

	select {
	case res := <-c0:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("waited too long for result - timeout 1")
	}

	c1 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 2"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("waited too long for result - timeout 2")
	}
}
