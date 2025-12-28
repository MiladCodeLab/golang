package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"runtime"
)

func main() {
	urls := []string{
		"http://www.google0.com",
		"http://www.golang0.org",
		"http://www.somestupidname0.com",
	}

	var g errgroup.Group
	g.SetLimit(3)

	for _, url := range urls {
		g.Go(func() error {
			_, err := http.Get(url)
			if err == nil {
				fmt.Printf("Successfully fetched %s\n", url)
			}
			return fmt.Errorf("error fetching %s: %v", url, err)
		})
	}
	println("num: ", runtime.NumGoroutine())

	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	} else {
		fmt.Printf("Failed to fetch a URL: %v\n", err)
	}
}
