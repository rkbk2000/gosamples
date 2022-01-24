package main

import (
	"fmt"

	"github.com/jinzhu/copier"
)

type Cat struct {
	age     int
	name    string
	friends []string
}

func checkCopyNotworkingAsExpected() {
	wilson := Cat{7, "Wilson", []string{"Tom", "Tabata", "Willie"}}
	nikita := Cat{}
	copier.CopyWithOption(&nikita, &wilson, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	fmt.Println(nikita)
	nikita.friends = append(nikita.friends, "Syd")

	fmt.Println(wilson)
	fmt.Println(nikita)
}
