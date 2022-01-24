package main

import (
	"fmt"
	"sync"
)

func lockDemo() {
	str := "A"

	c1 := make(chan int)

	var l sync.Mutex
	go func(c chan int) {
		l.Lock()
		str = "B"
		fmt.Println(str)
		l.Unlock()
		c <- 0
	}(c1)

	c2 := make(chan int)
	go func(c chan int) {
		l.Lock()
		str = "C"
		fmt.Println(str)
		l.Unlock()
		c <- 1
	}(c2)

	// Wait for the method completion
	<-c1
	<-c2
}
