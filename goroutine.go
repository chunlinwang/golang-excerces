package main

import (
	"fmt"
	"time"
)

func worker(finished chan bool) {
	fmt.Println("Worker: Started")
	time.Sleep(time.Second)
	fmt.Println("Worker: Finished")
	finished <- true
}

func main() {
	finished := make(chan bool)

	fmt.Println("Main: Starting worker")
	go worker(finished)

	fmt.Println("Main: Waiting for worker to finish")
	<-finished
	fmt.Println("Main: Completed")
}
