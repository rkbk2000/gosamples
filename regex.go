package main

import (
	"log"
	"regexp"
	"strings"
)

var monitorRegEx = regexp.MustCompile("monitor:([a-z0-9]*)")

func matches(conf string) bool {
	if !(monitorRegEx.MatchString(conf) && !(strings.HasSuffix(conf, "selfHealth") || strings.HasSuffix(conf, "resourceList"))) {
		log.Println(conf + " skipping")
		return true
	}
	log.Println(conf + " handling")
	return false
}
