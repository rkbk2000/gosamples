package main

import (
        "fmt"
        "sync"
)

// raceDemo demonstrates the race conditions between goroutines
// Compile with "-race" option to get the race warning
func raceDemo() {
        str := "A"

        wg := sync.WaitGroup{}
        wg.Add(1)

        go func(wsg *sync.WaitGroup) {
                defer wsg.Done()
                str = "B"
                fmt.Println(str)
        }(&wg)

        wg.Add(1)
        go func(wsg *sync.WaitGroup) {
                defer wsg.Done()
                str = "C"
                fmt.Println(str)
        }(&wg)

        wg.Wait()
}