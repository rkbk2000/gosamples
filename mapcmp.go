package main

import "fmt"

type Msg struct {
	URL                string
	Headers            map[string]string
	Body               []byte
	RetryCnt           int
	RetryIntervalInSec int
	//StatusCodes holds the status codes specified by the component to retry the rest call if it is failed
	StatusCodes map[int]bool
}

func handleMsg(msg *Msg) {
	clone := *msg
	clone.Headers = make(map[string]string)
	clone.Headers["h1"] = "v1"
	fmt.Printf("%v\n", clone)
}
