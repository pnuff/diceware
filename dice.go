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

	return 1 + int(r.Int64()), err
}

// Read the wordlist file in, map the dice roll values to the words on each line.
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

// ThrowDice rolls the given number of dice and returns them in an int slice.
func ThrowDice(numDice int) []int {
	var results = []int{}
	for i := 0; i < numDice; i++ {
		r, err := SixSided()
		if err != nil {
			fmt.Printf("Problem creating a random number.\n")
			os.Exit(1)
		}
		results = append(results, r)
	}
	return results
}

func main() {
	words := ReadWordlist(os.Stdin)
	var selected = []string{}
	for i := 0; i < 5; i++ {
		var dieRoll = []string{}
		results := ThrowDice(5)
		for _, r := range results {
			dieRoll = append(dieRoll, strconv.Itoa(r))
		}
		line := strings.Join(dieRoll, "")
		// TODO check for the line existing first
		selected = append(selected, words[line])
	}
	for _, s := range selected {
		fmt.Printf("%s ", s)
	}
	fmt.Printf("\n")
}
