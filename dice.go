package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// SixSided returns a random value between 1 and 6
func SixSided() int {
	r, err := rand.Int(rand.Reader, big.NewInt(int64(6)))
	if err != nil {
		return 0
	}
	return 1+int(r.Int64())
}


func main() {
	var results []int
	for i := 0; i < 5; i++ {
		r := SixSided()
		if r == 0 {
			fmt.Printf("Error creating a random number\n")
		}
		if r > 6 {
			fmt.Printf("\n\nError, roll is > 6: %d\n\n", r)
		}
		fmt.Printf("%d\t", r)
		results = append(results, r)
	}
	fmt.Printf("\ndone\n")

	for _, r := range results {
		fmt.Printf("%d", r)
	}
	fmt.Printf("\n")
}
