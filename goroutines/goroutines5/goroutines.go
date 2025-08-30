// https://reliasoftware.com/blog/concurrency-in-golang
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, ch chan string) {
	time.Sleep(time.Second)
	ch <- fmt.Sprintf("worker %d done", id)
}

func main() {
	data := make(chan string)
	wg := sync.WaitGroup{}
	for i := range 3 {
		wg.Add(1)
		go worker(i, data)
	}

	go func() {
		wg.Wait()
		close(data)
	}()

	for rcv := range data {
		fmt.Println(rcv)
	}
}
