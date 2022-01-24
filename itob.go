package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

// func GetBytes_(key interface{}) ([]byte, error) {

func iToB() {
	type rect struct {
		Width, Height int
	}

	r := rect{10, 20}

	var ir interface{}

	ir = r

	b, e := GetBytes(ir)

	if e != nil {
		log.Println(e)
	} else {
		fmt.Println("Rectangle1 =" + string(b))
	}
}

// GetBytes convets interface to bytes
func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil

}
