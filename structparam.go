package main

import "fmt"

type zInfo struct {
	name  string
	value int
}

func setValue(pzi *zInfo) {
	// pzi = &zInfo{"n1", 1}
	pzi.name = "n1"
}

func passByValue(zi zInfo) {
	setValue(&zi)
}

func checkPassByValue() {
	zi := zInfo{}
	passByValue(zi)
	fmt.Printf("zi = %v \n", zi)
}
