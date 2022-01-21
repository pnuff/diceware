package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io"
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

// Read the wordlist file in, map the dice roll values to the words on each line
func ReadWordlist(r io.Reader) map[string]string {
	var wordList = make(map[string]string)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		c := scanner.Text()
		temp := strings.Fields(c)
		// TODO check for incorrect number of fields
		wordList[temp[0]] = temp[1]
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}
	return wordList
}

func main() {
	words := ReadWordlist(os.Stdin)
	fmt.Printf("%+v\n", words)
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
