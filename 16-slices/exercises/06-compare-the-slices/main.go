// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"sort"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Compare the slices
//
//  1. Split the namesA string and get a slice
//
//  2. Sort all the slices
//
//  3. Compare whether the slices are equal or not
//
//
// EXPECTED OUTPUT
//
//   They are equal.
//
//
// HINTS
//
//   1. strings.Split function splits a string and
//      returns a string slice
//
//   2. Comparing slices: First check whether their length
//      are the same or not; only then compare them.
//
// ---------------------------------------------------------

func main() {
	const (
		equalMsg     = "They are equal"
		differentMsg = "They are NOT equal"
	)
	namesA := "Da Vinci, Wozniak, Carmack"
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}
	namesASlice := strings.Split(namesA, ", ")

	sort.Strings(namesASlice)
	sort.Strings(namesB)

	if len(namesASlice) == len(namesB) {
		equal := false
	mainLoop:
		for _, na := range namesASlice {
			for _, nb := range namesB {
				if na == nb {
					equal = true
					continue mainLoop
				} else {
					equal = false
				}
			}
		}
		if equal {
			fmt.Println(equalMsg)
		} else {
			fmt.Println(differentMsg)
		}
	}
}
