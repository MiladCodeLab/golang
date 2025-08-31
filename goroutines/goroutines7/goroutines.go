// https://goplay.tools/snippet/TyZVdIqfL5x
package main

import (
	"fmt"
	"time"
)

func main() {

	select {
	case <-time.After(3 * time.Second):
		for {
			select {
			case t, e := <-time.Tick(time.Second):
				fmt.Println("tick ", t, " ", e)

			}
		}
	}
}
