// https://reliasoftware.com/blog/concurrency-in-golang
package main

import "time"

func main() {
	done := make(chan struct{})

	go func() {
		i := 0
		for range time.Tick(time.Second) {
			i++
			if i > 10 {
				done <- struct{}{}
				break
			}
			println(i, " second")
		}
	}()
	<-done
}
