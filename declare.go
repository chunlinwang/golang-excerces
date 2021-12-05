package main

import (
	"fmt"
	"time"
)

func main() {
	var number int

	number = 10

	now := time.Now()

	fmt.Println("Variable number: ", number)
	fmt.Println("Variable now:", now)
}
