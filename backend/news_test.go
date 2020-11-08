package main

import "testing"

func TestFunc(t *testing.T) {
	var x int = 4
	if x == 6 {
		t.Errorf("error")
	}
}
