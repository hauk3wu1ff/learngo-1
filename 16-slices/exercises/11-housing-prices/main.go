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
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Housing Prices
//
//  We have received housing prices. Your task is to load the data into
//  appropriately typed slices then print them.
//
//  1. Check out the expected output
//
//
//  2. Check out the code below
//
//     1. header   : stores the column headers
//     2. data     : stores the real data related to each column
//     3. separator: you will use it to separate the data
//
//
//  3. Parse the data
//
//     1. Separate it into rows by using the newline character ("\n")
//
//     2. For each row, separate it by using the separator (",")
//
//
//  4. Load the data into distinct slices
//
//     1. Load the locations into a []string
//     2. Load the sizes into []int
//     3. Load the beds into []int
//     4. Load the baths into []int
//     5. Load the prices into []int
//
//
//  5. Print the header
//
//     1. Separate it by using the separator
//
//     2. Print each column 15 character wide ("%-15s")
//
//
//  6. Print the rows from the slices that you've created, line by line
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//
// HINTS
//
//  + strings.Split function can separate a string into a []string
//
// ---------------------------------------------------------

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	var (
		locations []string
		size      []int
		beds      []int
		baths     []int
		prices    []int
	)

	hs := strings.Split(header, separator)
	ds := strings.Split(data, "\n")

	for r, rv := range ds {
		for c, fv := range strings.Split(rv, ",") {
			switch c {
			case 0:
				locations = append(locations, fv)
			case 1:
				i, err := strconv.Atoi(fv)
				if err != nil {
					fmt.Printf("Row/Column %d/%d should be int\n", r, c)
					return
				}
				size = append(size, i)
			case 2:
				i, err := strconv.Atoi(fv)
				if err != nil {
					fmt.Printf("Row/Column %d/%d should be int\n", r, c)
					return
				}
				beds = append(beds, i)
			case 3:
				i, err := strconv.Atoi(fv)
				if err != nil {
					fmt.Printf("Row/Column %d/%d should be int\n", r, c)
					return
				}
				baths = append(baths, i)
			case 4:
				i, err := strconv.Atoi(fv)
				if err != nil {
					fmt.Printf("Row/Column %d/%d should be int\n", r, c)
					return
				}
				prices = append(prices, i)
			}
		}
	}
	fmt.Printf("%-15s%-15s%-15s%-15s%-15s\n", hs[0], hs[1], hs[2], hs[3], hs[4])
	for i := 1; i <= 75; i++ {
		fmt.Printf("%s", "=")
	}
	fmt.Println()
	for i := 0; i < len(ds); i++ {
		fmt.Printf("%-15s", locations[i])
		fmt.Printf("%-15d", size[i])
		fmt.Printf("%-15d", beds[i])
		fmt.Printf("%-15d", baths[i])
		fmt.Printf("%-15d\n", prices[i])
	}
}
