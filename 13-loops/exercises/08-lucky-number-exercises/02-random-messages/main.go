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
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
//----------------------------------------------------------
// Decision: I'm not using a switch statement.
// Instead I directly use the random int as an index into the
// appropriate winner or looser message array.
// ---------------------------------------------------------

const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your number is, harder it gets.

Wanna play?
`
)

var (
	winnerMsg [4]string = [4]string{"🎉  YOU WON!", "YOU'RE AWESOME", "ANOTHER WINNER!",
		"YOU WIN AGAIN, THIS TOO EASY FOR YOU!"}
	loserMsg [4]string = [4]string{"LOSER!", "YOU LOST", "THAT IS A LOSS", "YOU LOOSE!"}
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess <= 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	for turn := 1; turn <= maxTurns; turn++ {
		n := rand.Intn(guess) + 1

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

			/*
			 Choose a winner message at random.
			 Please note regarding rand.Intn: It returns a pseudo-random number in the
			 interval [0,n). The internal border excludes n.
			*/
			fmt.Printf("🎉  %s\n", winnerMsg[rand.Intn(len(winnerMsg)+1)])
		}
		return
	}
	// Choose a loser message at random.
	fmt.Printf("☠️  %s... Try again?\n", loserMsg[rand.Intn(len(loserMsg)+1)])
}
