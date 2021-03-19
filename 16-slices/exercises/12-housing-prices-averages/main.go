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
// EXERCISE: Housing Prices and Averages
//
//  Use the previous exercise to solve this exercise (Housing Prices).
//
//  Your task is to print the averages of the sizes, beds, baths, and prices.
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
//  ===========================================================================
//                 237.50         4.75           2.75           425000.00
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
	fmt.Printf("%s\n", strings.Repeat("=", 75))
	for i := 0; i < len(ds); i++ {
		fmt.Printf("%-15s", locations[i])
		fmt.Printf("%-15d", size[i])
		fmt.Printf("%-15d", beds[i])
		fmt.Printf("%-15d", baths[i])
		fmt.Printf("%-15d\n", prices[i])
	}
	fmt.Printf("%s\n", strings.Repeat("=", 75))
	fmt.Printf("%-15s%-15.2f%-15.2f%-15.2f%-15.2f\n", "", average(size), average(beds), average(baths), average(prices))
}

// Calculates the average of an int slice.
// Returns 0 for a nil or empty slice
func average(s []int) float64 {
	var sum int
	if len(s) > 0 {
		for _, v := range s {
			sum += v
		}
		return float64(sum) / float64(len(s))
	} else {
		return 0
	}
}
