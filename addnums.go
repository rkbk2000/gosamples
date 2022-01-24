package main

import (
	"fmt"
	"os"
)

// Wait for a number from user
func readNum(ic chan<- int, cont <-chan bool) {
	for {
		var val int
		fmt.Printf("Enter a number to add(-1 to exit): ")
		fmt.Fscanf(os.Stdin, "%d\n", &val)
		fmt.Println()
		ic <- val
		// Wait until the sum is printed
		<-cont
	}
}

func printSum() {
	ic := make(chan int)
	cont := make(chan bool)
	sum := 0

	go func() {
		readNum(ic, cont)
	}()

	for val := range ic {
		if val != -1 {
			sum += val
			fmt.Println("sum = ", sum)
			cont <- true
			continue
		}
		break
	}
	close(ic)
	close(cont)
}
