package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, name string) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s got context err: %s\n", name, ctx.Err())
			fmt.Printf("%s got context cause: %s\n", name, context.Cause(ctx))
			return
		default:
			fmt.Printf("%s did some work\n", name)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup

	ctx0, cancel0 := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx0, &wg, "worker-with-cancel")

	go func() {
		time.Sleep(3 * time.Second)
		cancel0()
	}()

	/////////////////////////

	// showcase Cause addition from Go 1.20
	ctx1, cancel1 := context.WithCancelCause(context.Background())
	wg.Add(1)
	go worker(ctx1, &wg, "worker-with-cancel-cause")

	go func() {
		time.Sleep(5 * time.Second)
		cancel1(errors.New("i just wanted the worker to stop"))
	}()

	/////////////////////////

	ctx2, _ := context.WithTimeoutCause(
		context.Background(), 6*time.Second, errors.New("the operation took too long"))
	wg.Add(1)
	go worker(ctx2, &wg, "worker-with-timeout-cause")

	wg.Wait()
	fmt.Println("\nall workers stopped")
}
