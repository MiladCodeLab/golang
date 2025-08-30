// https://reliasoftware.com/blog/concurrency-in-golang
package main

import (
	"math/rand"
	"time"
)

func main() {
	data := make(chan int, 2)

	data <- rand.Int() % 1000
	time.Sleep(200 * time.Millisecond) // simulate computation
	data <- rand.Int() % 1000
	//time.Sleep(200 * time.Millisecond) // simulate computation
	//data <- rand.Int() % 1000

	val1, ok := <-data
	println(val1, " ", ok)
	val2, ok := <-data
	println(val2, " ", ok)
	//val3, ok := <-data
	//println(val3, " ", ok)
}
