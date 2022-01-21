package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

// SixSided returns a random value between 1 and 6
func SixSided() (int, error) {
	r, err := rand.Int(rand.Reader, big.NewInt(int64(6)))
	if err != nil {
		fmt.Printf("Error calling rand.Int\n")
		return 0, err
	}
	if int(r.Int64()) < 0 || int(r.Int64()) > 5 {
		fmt.Printf("Error: die has unexpected number of sides\n")
		return 0, err
	}

	return 1+int(r.Int64()), err
}


func main() {
	var results, dieRoll = []int{}, []string{}
	for i := 0; i < 5; i++ {
		r, err := SixSided()
		if err != nil {
			fmt.Printf("Problem creating a random number.\n")
			os.Exit(1)
		}
		results = append(results, r)
	}

	for _, r := range results {
		dieRoll = append(dieRoll, strconv.Itoa(r))
	}
	line := strings.Join(dieRoll, "")
	fmt.Printf("%s\n", line)
}
