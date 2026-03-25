package main

import (
	"errors"
)

// Checking the behavior of if along with error check

func getErr(v int) error {
	if v%2 == 0 {
		return errors.New("even")
	}
	return nil
}

func checkIfErrorIf(err error, val int) {

	// This does not work: if err!=nil;err = getErr(val); err != nil {
	if err = getErr(val); err != nil {
	}
}
