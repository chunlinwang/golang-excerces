package main

import (
	"fmt"
)

type MyError struct{}

func (m *MyError) Error() string {
	return "error"
}

func div(a int, b int) (int, error) {
	if b == 0 {
		return 0, &MyError{}
	}

	return a / b, nil
}

func main() {
	r, err := div(10, 0)
	if err != nil {
		fmt.Println("Le dénominateur ne peut pas être zéro")

		return
	}

	fmt.Println(r)
}
