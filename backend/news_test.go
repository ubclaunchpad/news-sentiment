package main

import "testing"

func TestFunc(t *testing.T) {
	if 6 != 6 {
		t.Errorf("error")
	}
}
