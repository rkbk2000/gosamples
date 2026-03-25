package main

import "fmt"

type MyStruct struct {
	data int
}

// Function by Value is a problem as the method would get called on the object that
// it was called on. Always use pointer if you want to reassign values to the same object.
// func (ms *MyStruct) Close()
func (ms MyStruct) Close() {
	fmt.Printf("In my struct close, data = %d", ms.data)
}

func deferTest() {
	m1 := MyStruct{10}
	defer m1.Close()

	// Assign a new object
	m1 = MyStruct{22}
}
