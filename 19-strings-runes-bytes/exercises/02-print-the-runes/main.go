// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print the runes
//
//  1. Loop over the "console" word and print its runes one by one,
//     in decimals, hexadecimals and binary.
//
//  2. Manually put the runes of the "console" word to a byte slice, one by one.
//
//     As the elements of the byte slice use only the rune literals.
//
//     Print the byte slice.
//
//  3. Repeat the step 2 but this time, as the elements of the byte slice,
//     use only decimal numbers.
//
//  4. Repeat the step 2 but this time, as the elements of the byte slice,
//     use only hexadecimal numbers.
//
//
// EXPECTED OUTPUT
//   Run the solution to see the expected output.
// ---------------------------------------------------------

func main() {
	const word = "console"

	var (
		rbword []byte
		dbword []byte
		hbword []byte
	)
	for _, rune := range word {
		// 1) print its runes one by one, as unicode-point, decimals, hexadecimals and binary.
		fmt.Printf("%[1]U, %[1]d, %[1]x, %[1]b\n", rune)
		// 2) Manually put the runes of the "console" word to a byte slice, one by one.
		rbword = append(rbword, byte(rune))
	}

	// Print the byte slice.
	fmt.Printf("with runes: %s\n", rbword)
}
