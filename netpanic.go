package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func netpanic() {
	r := strings.NewReader("<template><tBody><isindex/action=0>")
	html.Parse(r)
	r = strings.NewReader("<math><template><mo><template>")
	html.Parse(r)
	fmt.Println("No panic")
}
