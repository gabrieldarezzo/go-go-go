package main

import "testing"

func TestChange(t *testing.T) {
	change := RetrieveChange(55)

	if change != 2 {
		t.Errorf("Change is invalid")
	}

}
