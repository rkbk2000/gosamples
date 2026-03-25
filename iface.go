package main

import "fmt"

// Demonstration of interfaces

// Hold multiple different types

type Object struct {
	Value string
}

type MetricData interface {
}

type t1 struct {
	val string
}

type t2 struct {
	val1 string
	val2 string
}

var cache []interface{}

func printAll() {
	mixedArray := []interface{}{"astring", 10, &Object{"hello"}}
	for i, v := range mixedArray {
		fmt.Printf("Value[%d] = %v ", i, v)
	}

	fmt.Println()
	cache = append(cache, t1{"a"})
	cache = append(cache, t2{"b", "c"})
	for i, v := range cache {
		fmt.Printf("Value[%d] = %v ", i, v)
	}
	fmt.Println()
}
