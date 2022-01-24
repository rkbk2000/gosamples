package main

import (
	"fmt"
	"time"
)

var myval int

func startTicker(f func()) chan bool {
	done := make(chan bool, 1)
	go func() {
		ticker := time.NewTicker(time.Second * 1)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				myval++
				f()
				time.Sleep(time.Second * 1)
			case <-done:
				myval++
				fmt.Println("done")
				return
			}
		}
	}()
	return done
}

func checkTicker() {
	done := startTicker(func() {
		fmt.Println("tick...")
	})
	time.Sleep(5 * time.Second)
	close(done)
	time.Sleep(5 * time.Second)
}
