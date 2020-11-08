package main

import "testing"

func TestServer(t *testing.T) {
	error := runServer()
	if error != nil {
		t.Errorf("error")
	}
}
