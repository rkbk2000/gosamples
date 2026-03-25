package main

import (
	"fmt"
	"runtime"
)

// printMemoryUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func printMemoryUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// See: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// Checks memory usage for a count of strings with an average size of elSize
func checkMemForStrings(elSize, count int) {
	fmt.Printf("Memroy usage before for element size = %d, count = %d\n", elSize, count)
	printMemoryUsage()

	str := ""
	for i := 0; i < elSize; i++ {
		str += string(byte(i + 'A'))
	}
	var arr []string
	for i := 0; i < elSize; i++ {
		arr = append(arr, str)
	}
	fmt.Printf("Memroy usage after allocation for element size = %d, count = %d\n", elSize, count)
	printMemoryUsage()
}

// Checks memory usage for a count of ints with an average size of elSize
func checkMemForInts(count int) {
	fmt.Printf("Memroy usage before for element count = %d\n", count)
	printMemoryUsage()

	var arr []int
	for i := 0; i < count; i++ {
		arr = append(arr, i)
	}
	fmt.Printf("Memroy usage after allocation for element count = %d\n", count)
	printMemoryUsage()
}
