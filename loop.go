package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()

	n := 0

	for i := 0; i < 1000000000; i++ {
		n += i
	}

	t2 := time.Now()

	fmt.Println("t1 : ", t1.UnixMilli())

	fmt.Println("t2 : ", t2.UnixMilli())

	fmt.Println("time : ", t2.UnixMilli()-t1.UnixMilli())
}
