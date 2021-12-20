package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	fmt.Printf("Worker %v: Started at %v\n", id, time.Now())
	time.Sleep(time.Second)
	fmt.Printf("Worker %v: Finished at %v\n", id, time.Now())
}

func main() {
	var wg sync.WaitGroup

	t1 := time.Now()
	fmt.Println("Main: Start at : ", t1)

	for i := 0; i < 5; i++ {
		fmt.Println("Main: Starting worker", i, time.Now())
		wg.Add(1)
		go worker(&wg, i)
	}

	fmt.Println("Main: Waiting for workers to finish", time.Now())
	wg.Wait()
	t2 := time.Now()
	fmt.Println("Main: Stop at : ", t2)
	fmt.Println("time execute : ", t2.UnixMilli()-t1.UnixMilli())
}
