package main

import (
	"testing"
)

// BenchmarkSixSided
func BenchmarkSixSided(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SixSided()
	}
}

// TestSixSidedBoundaries verifies SixSided returns between 1 and 6, inclusive.
func TestSixSidedBoundaries(t *testing.T) {
	r := SixSided()
	if r <= 0 || r > 6 {
		t.Fatalf("SixSided() returned a value out of bounds: %d\n", r)
	}
}
