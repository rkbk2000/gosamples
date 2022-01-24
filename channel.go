package main

import (
	"fmt"
	"time"
)

// channelDemo demonstrates the use of channels
func channelDemo() {
	str := "A"

	c1 := make(chan int)

	go func(c chan int) {
		fmt.Println(str)
		//time.Sleep(900 * time.Millisecond)
		c <- 0
	}(c1)

	str = "C"
	c2 := make(chan int)
	go func(c chan int) {
		fmt.Println(str)
		// time.Sleep(900 * time.Millisecond)
		c <- 1
	}(c2)

	// Wait for the goroutine completion
	<-c1
	<-c2
}

type configDetails struct {
	data               string
	discoveryCompleted chan bool
}

func runTogglingChannel(d string) {

	cd := &configDetails{
		data:               d,
		discoveryCompleted: make(chan bool),
	}
	go getResults(cd)

	// Wake up every second and check
	ticker(time.Duration(1), cd)

	for {
		time.Sleep(1 * time.Hour)
	}
}

func ticker(interval time.Duration, cd *configDetails) {
	ticker := time.NewTicker(interval * time.Second)

	go func() {
		toggleLoop(ticker, cd)
	}()
}

func getResults(cd *configDetails) {
	// sleep for 5 second
	time.Sleep(5 * time.Second)
	cd.discoveryCompleted <- true
}

func toggleLoop(ticker *time.Ticker, cd *configDetails) {
	defer ticker.Stop()
	for {
		select {
		case tm := <-ticker.C:
			fmt.Printf("Checking at %v\n", tm)
			select {
			case done := <-cd.discoveryCompleted:
				fmt.Printf("Previous task done, running next task %v\n", done)
				go getResults(cd)
			default:
				fmt.Println("Previous task still running, skipping")
			}
		}
	}
}
