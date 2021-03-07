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
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
// ---------------------------------------------------------

const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your number is, harder it gets.

Wanna play?
Use: go run main.go 4
For verbose output use: go run main.go -v 4
`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	var (
		guess   int
		err     error
		verbose bool
	)

	if len(args) == 1 {
		guess, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Not a number.")
			return
		}
		if guess <= 0 {
			fmt.Println("Please pick a positive number.")
			return
		}

	}

	if len(args) == 2 {
		if strings.ToLower(args[0]) == "-v" {
			verbose = true
		} else {
			fmt.Printf(usage, maxTurns)
			return
		}
		guess, err = strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Not a number.")
			return
		}
		if guess <= 0 {
			fmt.Println("Please pick a positive number.")
			return
		}
		//		fmt.Printf("Inside 'if len(args) == 2': Variable guess, type = %T, value = %[1]v\n", guess)

	}

	//	fmt.Printf("Variable guess, type = %T, value = %[1]v\n", guess)

	for turn := 1; turn <= maxTurns; turn++ {
		n := rand.Intn(guess) + 1
		if verbose {
			fmt.Printf("%d ", n)
		}

		// Better, why?
		//
		// Instead of nesting the if statement into
		//   another if statement; it simply continues.
		//
		// TLDR: Avoid nested statements.
		if n != guess {
			continue
		}

		if turn == 1 {
			fmt.Println("🥇 FIRST TIME WINNER!!!")
		} else {
			fmt.Println("🎉  YOU WON!")
		}
		return
	}

	fmt.Println("☠️  YOU LOST... Try again?")
}
