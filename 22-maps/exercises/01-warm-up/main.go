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
// EXERCISE: Warm-up
//
//  Create and print the following maps.
//
//  1. Phone numbers by last name
//  2. Product availability by Product ID
//  3. Multiple phone numbers by last name
//  4. Shopping basket by Customer ID
//
//     Each item in the shopping basket has a Product ID and
//     quantity. Through the map, you can tell:
//     "Mr. X has bought Y bananas"
//
// ---------------------------------------------------------

func main() {
	// Hint: Store phone numbers as text

	// #1
	// Key        : Last name
	// Element    : Phone number
	phoneDict := map[string]string{"Wulff": "+353 51 967925",
		"O'Fitzmahoney": "+353 51 12345",
		"Yeats":         "+353 51 67890"}
	fmt.Printf("The phone number dictionary %#v\n", phoneDict)

	// #2
	// Key        : Product ID
	// Element    : Available / Unavailable
	productDict := map[string]bool{
		"mon00001": true,
		"mon00002": false,
		"mon00003": true,
	}
	fmt.Println("The product ID dictionary")
	for k, v := range productDict {
		if v {
			fmt.Printf("ProductId %s is in stock\n", k)
		} else {
			fmt.Printf("ProductId %s is not in stock\n", k)
		}
	}

	// #3
	// Key        : Last name
	// Element    : Phone numbers
	mpDict := map[string]map[string]string{
		"Wulff":         {"home": "+353 51 967925", "mobile": "+353 858779400"},
		"O'Fitzmahoney": {"home": "+353 51 123456", "mobile": "+353 851234567"},
		"Yeats":         {"home": "+353 51 678901", "mobile": "+353 852345678"},
	}
	fmt.Printf("The multi phone number dictionary %#v\n", mpDict)

	// #4
	// Key        : Customer ID
	// Element Key:
	//   Key: Product ID Element: Quantity
	basketDict := map[int]map[int]int{
		1000001: {1234567: 12, 23456789: 1, 34567890: 2},
		1000002: {1234567: 22, 23456789: 2, 34567890: 4},
		1000003: {1234567: 32, 23456789: 3, 34567890: 5},
	}
	fmt.Printf("The shopping basket dictionary %#v\n", basketDict)
}
