package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

//const workerNum =3
//type WorkerPool struct {
//	input chan
//}

func main() {
	file1, err := os.Open("file1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()

	file2, err := os.Open("file2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()

	file3, err := os.Open("file3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file3.Close()
	files := []*os.File{file1, file2, file3}
	input := make(chan string, 3)
	result := make(map[string]int)
	wg := sync.WaitGroup{}

	// Read segment
	for id := range 3 {
		wg.Go(func() {
			log.Printf("worker id %d\n", id)
			f := files[id]
			reader := bufio.NewReader(f)
			for line, _, err := reader.ReadLine(); err != io.EOF; line, _, err = reader.ReadLine() {
				log.Printf("line %d: %s\n", id, line)
				input <- string(line)
			}

		})
	}
	go func() {
		wg.Wait()
		close(input)
	}()
	// Process segment
	for line := range input {
		l := strings.ToLower(line)
		t := strings.ReplaceAll(l, ",", "")
		words := strings.Split(t, " ")
		for _, word := range words {
			w := strings.TrimSpace(word)
			val, ok := result[w]
			if ok {
				result[word] = val + 1
			} else {
				result[word] = 1
			}
		}
	}

	fmt.Printf("result is %v\n", result)
}
