package main

import "testing"

func TestMain(m *testing.M) {
	// set up the global map
	flips = getFlipMap()
}
