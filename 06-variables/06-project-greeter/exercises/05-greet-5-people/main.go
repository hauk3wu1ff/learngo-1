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
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Greet 5 People
//
//  Greet 5 people this time.
//
//  Please do not copy paste from the previous exercise!
//
// RESTRICTION
//  This time do not use variables.
//
// INPUT
//  bilbo balbo bungo gandalf legolas
//
// EXPECTED OUTPUT
//  There are 5 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Hello great gandalf !
//  Hello great legolas !
//  Nice to meet you all.
// ---------------------------------------------------------

func main() {
	// TYPE YOUR CODE HERE
	if len(os.Args) < 6 {
		fmt.Println("greeter5 greets exactly 5 persons, usage: ./greeter5 p1 p2 p3, p4, p5")
		os.Exit(1)
	}
	var (
		l       = len(os.Args) - 1
		person1 = os.Args[1]
		person2 = os.Args[2]
		person3 = os.Args[3]
		person4 = os.Args[4]
		person5 = os.Args[5]
	)

	fmt.Println("There are", l, "persons.")
	fmt.Println("Hello great", person1, "!")
	fmt.Println("Hello great", person2, "!")
	fmt.Println("Hello great", person3, "!")
	fmt.Println("Hello great", person4, "!")
	fmt.Println("Hello great", person5, "!")
	fmt.Println("Nice to meet you all.")
}
