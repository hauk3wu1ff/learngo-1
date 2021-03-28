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

	"github.com/inancgumus/screen"
	"github.com/mattn/go-runewidth"
)

// ---------------------------------------------------------
// EXERCISE: Previous positions
//
//  Let's optimize the program once more. This time you're
//  going to optimize the clearing off the previous positions.
//
//  1. Find the code below marked as "remove the previous ball"
//
//  2. Instead of clearing every position on the board to false,
//     only set the previous position to false. So, don't use
//     a loop, remove it.
//
//  3. Change the velocity of the ball like so:
//
//     vx, vy = 5, 2
//
//  4. Run the program and solve the problem
//
//
// HINT
//
//  Don't forget saving the previous position.
//
// ---------------------------------------------------------

func main() {
	const (
		cellEmpty = ' '
		cellBall  = '⚾'

		maxFrames = 1200
		speed     = time.Second / 20
	)

	var (
		px, py   int    // ball position
		opx, opy int    // previous ball position
		vx, vy   = 5, 2 // velocities

		cell rune // current cell (for caching)
	)

	// you can get the width and height using the screen package easily:
	width, height := screen.Size()

	// get the rune width of the ball emoji
	ballWidth := runewidth.RuneWidth(cellBall)

	// adjust the width and height
	width /= ballWidth
	height-- // there is a 1 pixel border in my terminal

	// create the board
	board := make([][]bool, width)
	for column := range board {
		board[column] = make([]bool, height)
	}

	// drawing buffer length
	// *2 for extra spaces
	// +1 for newlines
	bufLen := (width*2 + 1) * height

	// create a drawing buffer
	buf := make([]rune, 0, bufLen)

	// clear the screen once
	screen.Clear()

	for i := 0; i < maxFrames; i++ {
		// calculate the next ball position
		opx, opy = px, py // previous ball position for next iteration
		px += vx
		py += vy

		// ensure px, py are not 'index out of range'
		if px <= 0 {
			px = 0
		}
		if px >= width-1 {
			px = width - 1
		}
		if py <= 0 {
			py = 0
		}
		if py >= height-1 {
			py = height - 1
		}

		// when the ball hits a border reverse its direction
		if px <= 0 || px >= width-1 {
			vx *= -1
		}
		if py <= 0 || py >= height-1 {
			vy *= -1
		}

		// remove the previous ball
		// for y := range board[0] {
		// 	for x := range board {
		// 		board[x][y] = false
		// 	}
		// }
		board[opx][opy] = false

		// put the new ball
		board[px][py] = true

		// rewind the buffer (allow appending from the beginning)
		buf = buf[:0]

		// draw the board into the buffer
		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
				}
				buf = append(buf, cell, ' ')
			}
			buf = append(buf, '\n')
		}

		// print the buffer
		screen.MoveTopLeft()
		fmt.Print(string(buf))

		// slow down the animation
		time.Sleep(speed)
	}
}
