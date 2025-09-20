package main

import (
	"fmt"
	"time"
)

const WorkerNum int = 3

type JobFunc func()

type Worker struct {
	Input chan JobFunc
}

func (w *Worker) Start() {
	w.Input = make(chan JobFunc, WorkerNum)
	for i := range WorkerNum {
		go w.worker(i)
	}
}
func (w *Worker) worker(workerID int) {
	for job := range w.Input {
		println("worker ", workerID)
		job()
	}
}

func (w *Worker) Close() {
	close(w.Input)
}

func main() {
	//wg := sync.WaitGroup{}
	w := Worker{}
	w.Start()
	w.Input <- func() {
		time.Sleep(time.Millisecond * 170)
		fmt.Println("job1 is running")
	}

	w.Input <- func() {
		time.Sleep(time.Millisecond * 190)
		fmt.Println("job4 is running")
		fmt.Println("simple job to run")
	}

	w.Input <- func() {
		fmt.Println("job3 is running")
		time.Sleep(time.Millisecond * 130)
	}

	w.Input <- func() {
		fmt.Println("job2 is running")
		time.Sleep(time.Millisecond * 120)
	}
	time.Sleep(time.Second)
	defer w.Close()

}
