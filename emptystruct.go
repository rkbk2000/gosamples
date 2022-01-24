package main

import (
	"errors"
	"log"
)

// Response is a struct containing the code
type Response struct {
	code int
}

func getResponse() (Response, error) {
	return Response{}, errors.New("Empty response")
}

func checkResponse() {
	res, _ := getResponse()
	log.Println(res.code)
}
