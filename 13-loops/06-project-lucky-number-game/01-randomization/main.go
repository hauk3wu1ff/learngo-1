// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// rand.Seed(10)
	// rand.Seed(100)

	// t := time.Now()
	// rand.Seed(t.UnixNano())

	// ^-- same:
	t := time.Now().UnixNano()
	fmt.Println("The seed t (UnixNano) is", t)
	rand.Seed(t)

	// Read Int from Stdin
	guess := readInt()
	fmt.Printf("Pick is %d.\nTrying to guess this number...\n", guess)
	for n := -1; n != guess; {
		n = rand.Intn(guess + 1)
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}

// Read int from Stdin
func readInt() int {
	fmt.Print("Pick a number >= 0, please, any number: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	number := 42
	if scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading input:", err)
			os.Exit(1)
		}
		text := scanner.Text()
		guess, err := strconv.Atoi(text)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error %s converting %s to int\n", err, text)
			os.Exit(1)
		}
		number = guess
	}
	return number
}
