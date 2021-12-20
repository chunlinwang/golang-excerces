package object

import "fmt"

type PubObject struct {
	name string // private
	Age  int    // public
}

// public method
func Hello() string {
	return fmt.Sprintf("hello")
}

// private method
func hey() string {
	return fmt.Sprintf("hello")
}

// public
func (o *PubObject) GetName() string {
	return o.name
}

// private
func (o *PubObject) getage() int {
	return o.Age
}
