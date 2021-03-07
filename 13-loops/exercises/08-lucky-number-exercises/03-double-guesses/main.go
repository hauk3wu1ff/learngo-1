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
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------

const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one of those numbers with two guesses.

The greater your number is, harder it gets.

Wanna play?
`
)

var (
	winnerMsg [4]string = [4]string{"YOU WON!", "YOU'RE AWESOME", "ANOTHER WINNER!",
		"YOU WIN AGAIN, THIS TOO EASY FOR YOU!"}
	loserMsg [4]string = [4]string{"LOSER!", "YOU LOST", "THAT IS A LOSS", "YOU LOOSE!"}
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Printf(usage, maxTurns)
		return
	}
	var (
		guess [2]int
		err   error
		max   int
	)
	for i, v := range args {
		guess[i], err = strconv.Atoi(v)
		if err != nil {
			fmt.Println("Not a number.")
			return
		}

		if guess[i] <= 0 {
			fmt.Println("Please pick a positive number.")
			return
		}
		if guess[i] > max {
			max = guess[i]
		}
	}

	for turn := 1; turn <= maxTurns; turn++ {
		n := rand.Intn(max) + 1

		// Better, why?
		//
		// Instead of nesting the if statement into
		//   another if statement; it simply continues.
		//
		// TLDR: Avoid nested statements.
		if n != guess[0] && n != guess[1] {
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
			fmt.Println("len(winnerMsg) =", len(winnerMsg))
			fmt.Printf("🎉  %s\n", winnerMsg[rand.Intn(len(winnerMsg))])
		}
		return
	}
	// Choose a loser message at random.
	fmt.Printf("☠️  %s... Try again?\n", loserMsg[rand.Intn(len(loserMsg))])
}
