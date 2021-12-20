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

func workerInt(intVal chan int) {
	fmt.Println("Worker intVal: Started")
	time.Sleep(time.Second)
	fmt.Println("Worker intVal: Finished")
	intVal <- 999
}

func main() {
	finished := make(chan bool)

	fmt.Println("Main: Starting worker")
	go worker(finished)

	fmt.Println("Main: Waiting for worker to finish")
	data := <-finished

	fmt.Println("Main: Completed", data)

	intVal := make(chan int)
	go workerInt(intVal)

	intV := <-intVal

	fmt.Println("Main: intVal", intV)
}
