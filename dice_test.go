package dice

import (
	"testing"
)

func TestRandInt(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		min  int
		max  int
	}{
		{
			name: "TestRandInt",
			min:  12,
			max:  34,
		},
	}
	//
	for _, c := range tests {
		t.Run("", func(t *testing.T) {
			output := RandInt(c.min, c.max)

			// c.min<=output<=c.max
			assert.True(t, (c.min <= output && output <= c.max))
		})
	}
}

// BenchmarkRandInt
func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandInt(0, 99999)
	}
}
