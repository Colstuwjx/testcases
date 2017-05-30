package main

import (
	"context"
	"fmt"
	"sync"
)

func testSimple() {
	var wg sync.WaitGroup
	wg.Add(1)

	gen := func(ctx context.Context, wg *sync.WaitGroup) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("go rountine done.")
					wg.Done()
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	// defer cancel() // not really needed.

	// range gen simply mock up the complex query in the real world.
	for n := range gen(ctx, &wg) {
		fmt.Println(n)
		if n == 5 {
			cancel()
			break
		}
	}

	wg.Wait()
}

func main() {
	// TODO: spawn some workers to do some mock query,
	// and test whether Context Timeout/Cancel works as the expected behavior.
	testSimple()
}
