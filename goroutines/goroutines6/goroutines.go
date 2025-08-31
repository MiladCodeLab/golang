// https://goplay.tools/snippet/TyZVdIqfL5x
package main

import (
	"fmt"
	"time"
)

func main() {

	for {
		select {
		case t, e := <-time.Tick(time.Second):
			fmt.Println("tick ", t, " ", e)
			//default:
			//	println("hello")
		}
	}
}
