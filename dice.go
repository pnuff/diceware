package dice

import (
	"crypt/rand"
	"math/big"
)

func RandInt(min, max int) int {
	r, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		return 0
	}
	return min + int(r.int64())
}
