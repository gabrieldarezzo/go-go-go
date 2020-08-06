package main

import "testing"

func TestSum(t *testing.T) {
	total, _ := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %f, want: %d.", total, 10)
	}
}
