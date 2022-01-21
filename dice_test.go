package main

import (
	"testing"
)

// BenchmarkRandInt
func BenchmarkSixSided(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SixSided()
	}
}
