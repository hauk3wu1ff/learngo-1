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
	"time"
)

func main() {
	type placeholder [5]string

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	separator := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	ct := time.Now()
	h, m, s := ct.Hour(), ct.Minute(), ct.Second()
	fmt.Printf("%d:%d:%d\n", h, m, s)
	hHigh := h / 10
	hLow := h - hHigh*10
	mHigh := m / 10
	mLow := m - mHigh*10
	sHigh := s / 10
	sLow := s - sHigh*10
	fmt.Printf("%d%d:%d%d:%d%d\n", hHigh, hLow,
		mHigh, mLow, sHigh, sLow)

	clock := [8]placeholder{
		digits[hHigh],
		digits[hLow],
		separator,
		digits[mHigh],
		digits[mLow],
		separator,
		digits[sHigh],
		digits[sLow],
	}
	for line := range clock[0] {
		for _, digit := range clock {
			fmt.Print(digit[line], " ")
		}
		fmt.Println()
	}
}
