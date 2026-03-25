package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func get(h, p string) error {
	u := "http://" + h + ":" + p

	// IF you want to support the URLs correctly, use PathEscape
	// u := "http://" + url.PathEscape(h) + ":" + p

	fmt.Println("GET on url: ", u)

	r, e := http.Get(u)

	if e == nil {
		b, e := ioutil.ReadAll(r.Body)

		if e == nil {
			fmt.Println("Response = ", string(b))
		}
		return nil
	}
	fmt.Println("Error:", e.Error())
	fmt.Println()
	return e
}

func getUsingURL(h, p string) error {
	url := url.URL{Scheme: "http", Host: h + ":" + p}

	fmt.Println("GET using golang net URL package for url: ", url.String())

	r, e := http.Get(url.String())

	if e == nil {
		b, e := ioutil.ReadAll(r.Body)

		if e == nil {
			fmt.Println("Response = ", string(b))
		}
		return nil
	}
	fmt.Println("Error:", e.Error())
	fmt.Println()
	return e
}
