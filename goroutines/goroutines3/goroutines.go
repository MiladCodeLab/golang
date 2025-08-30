// https://reliasoftware.com/blog/concurrency-in-golang
package main

import (
	"math/rand"
	"time"
)

func main() {
	data := make(chan int)

	go func() {
		time.Sleep(1 * time.Second) // simulate computation
		data <- rand.Int() % 1000
	}()

	val, ok := <-data
	println(val, " ", ok)
}
