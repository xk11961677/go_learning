/**
接口
*/
package main

import "fmt"

// defined interface
type OrderCreator interface {
	valid()
	process()
	result()
}

// implement OrderCreator
type BusinessOrderCreator struct {
}

func (b BusinessOrderCreator) process() {
	fmt.Println("step 2")
}

func (b BusinessOrderCreator) result() {
	fmt.Println("step 2")
}

func (b BusinessOrderCreator) valid() {
	fmt.Println("step 1")
}

// main func
func createOrder(oc OrderCreator) {
	oc.valid()
	oc.valid()
	oc.result()
}

// entry point
func main() {
	createOrder(BusinessOrderCreator{})
}
