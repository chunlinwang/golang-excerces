package main

import (
	"fmt"

	"github.com/chunlinwang/golang-excerces/object"
)

func main() {
	pub := object.PubObject{
		Name: "Tom",
		Age:  34,
	}

	fmt.Println(pub.GetName())
	// 	fmt.Println(pub.getName()) error getName est private
}
