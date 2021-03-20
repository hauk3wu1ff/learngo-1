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
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Print daily requests
//
//  You've got request logs of a web server. The log data
//  contains 8-hourly totals per each day. It is stored
//  in the `reqs` slice.
//
//  Find and print the total requests per day, as well as
//  the grand total.
//
//  See the `reqs` slice and the steps in the code below.
//
//
// RESTRICTIONS
//
//   1. You need to produce the daily slice, don't just loop
//      and print the element totals directly. The goal is
//      gaining more experience in slice operations.
//
//   2. Your code should work even if you add to or remove the
//      existing elements from the `reqs` slice.
//
//      For example, after solving the exercise, try it with
//      this new data:
//
//      reqs := []int{
// 	      500, 600, 250,
// 	      200, 400, 50,
// 	      900, 800, 600,
// 	      750, 250, 100,
// 	      150, 654, 235,
// 	      320, 534, 765,
// 	      121, 876, 285,
// 	      543, 642,
// 	      // the last element is missing (your code should be able to handle this)
// 	      // that is why you shouldn't calculate the `size` below manually.
//      }
//
//      The grand total of the new data should be 10525.
//
//
// EXPECTED OUTPUT
//
//   Please run `solution/main.go` to see the expected
//   output.
//
//   go run solution/main.go
//
// ---------------------------------------------------------

func main() {
	// There are 3 request totals per day (8-hourly)
	const N = 3

	reqs := []int{
		500, 600, 250,
		200, 400, 50,
		900, 800, 600,
		750, 250, 100,
		150, 654, 235,
		320, 534, 765,
		121, 876, 285,
		543, 642,
	}
	// DAILY REQUESTS DATA (8-HOURLY TOTALS PER DAY)
	// reqs := []int{
	// 	500, 600, 250, // 1st day: 1350 requests
	// 	200, 400, 50, // 2nd day: 650 requests
	// 	900, 800, 600, // 3rd day: 2300 requests
	// 	750, 250, 100, // 4th day: 1100 requests
	// 	// grand total: 5400 requests
	// }

	// ================================================
	// #1: Make a new slice with the exact size needed.

	size := len(reqs) / N // you need to find the size.
	daily := make([][]int, 0, size)

	// ================================================

	// ================================================
	// #2: Group the `reqs` per day into the slice: `daily`.
	//
	// So the daily will be:
	// [
	//  [500, 600, 250]
	//  [200, 400, 50]
	//  [900, 800, 600]
	//  [750, 250, 100]
	// ]

	var d []int
	for i := 0; i < len(reqs); i = i + N {
		for j := 0; j < N; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
			fmt.Printf("(i + j) %% N = %d\n", (i+j)%N)
			if (i+j)%N == 0 {
				d = make([]int, 0, 3)
			}
			if i+j < len(reqs) {
				d = append(d, reqs[i+j])
			}
		}
		daily = append(daily, d)
	}
	// ================================================
	// #3: Print the results

	// Print a header
	fmt.Printf("%-10s%-10s\n", "Day", "Requests")
	fmt.Println(strings.Repeat("=", 20))

	// Loop over the daily slice and its inner slices to find
	// the daily totals and the grand total.
	for i, d := range daily {
		var total int
		for _, r := range d {
			fmt.Printf("%-10d%-10d\n", i, r)
			total += r
		}
		fmt.Printf("%9s %d\n", "TOTAL:", total)
	}
	// ================================================
}
