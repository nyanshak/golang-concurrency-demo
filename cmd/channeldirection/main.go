package main

import (
	"fmt"
)

/*
None of the following code will compile. This is here to demonstrate the
protections the compiler provides when using channel directions.
*/

func receiveOnSendOnlyChannel(ch chan<- int) {
	i := <-ch
	fmt.Println(i)
}

func sendOnReceiveOnlyChannel(ch <-chan int) {
	ch <- 1
}

func closeReceiveOnlyChannel(ch <-chan int) {
	close(ch)
}

func main() {
	// do nothing
}
