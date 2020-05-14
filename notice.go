package main

import (
	"math/rand"
	"runtime"
)

func main() {
	done := make(chan struct{})
	ch := generateInt(done)
	println(<-ch)
	println(<-ch)

	println(runtime.NumGoroutine())

	//停止生产
	close(done)

	println(<-ch)
	println(<-ch)

	println(runtime.NumGoroutine())
}

func generateInt(done chan struct{}) chan int {
	chanInt := make(chan int)

	go func() {
	Label:
		for {
			select {
			case chanInt <- rand.Int():
			case <-done:
				break Label
			}
		}

		close(chanInt)
	}()

	return chanInt
}
