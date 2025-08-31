package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func worker(input <-chan int, out chan<- int) {
	ti := rand.UintN(10000)
	fmt.Println("worker wait for ", ti, " ms")
	time.Sleep(time.Duration(ti) * time.Millisecond)
	n := <-input
	out <- n * 200
}
func main() {
	input := make(chan int)
	output := make(chan int)

	go func() {
		worker(input, output)
	}()

	go func() {
		input <- 41
	}()
	select {
	case <-time.After(time.Second):
		fmt.Println("timeout")
		return
	case val, ok := <-output:
		fmt.Println("main:", val, " ", ok)
	}

}
