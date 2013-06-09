package main

import "testing"

func TestKeys(t *testing.T) {
	const in, out = 2, 5
	if x := out - in; x != in {
		t.Errorf("%v - %v = %v; want %v", out, in, x, in)
	}
}
