package main

import (
	"testing"
)

func TestGetIPV6URL(t *testing.T) {
	tests := []struct {
		host        string
		port        string
		expectError bool
	}{{
		"fe80::9921:2933:14c8:b490%57", "80", true},
		{"fd80:9f3a:21e1:0:cdd0:98d1:3da3:ab58", "80", false},
		{"172.21.224.1", "80", false},
	}

	for _, tt := range tests {
		e := get(tt.host, tt.port)
		res := (e != nil)
		if tt.expectError != res {
			t.Error("For test:", tt.host, ":", tt.port, "expected ", tt.expectError, " found:", res)
		}
	}
}

func TestGetURLIPV6URL(t *testing.T) {
	tests := []struct {
		host        string
		port        string
		expectError bool
	}{{
		"fe80::9921:2933:14c8:b490%57", "80", false},
		{"fd80:9f3a:21e1:0:cdd0:98d1:3da3:ab58", "80", false},
		{"172.21.224.1", "80", false},
	}

	for _, tt := range tests {
		e := getUsingURL(tt.host, tt.port)
		res := (e != nil)
		if tt.expectError != res {
			t.Error("For test:", tt.host, ":", tt.port, "expected ", tt.expectError, " found:", res)
		}
	}
}

// Conclusion: We are safe if using golang's url.URL
// Else URL encoding needs to be done on URL strings
