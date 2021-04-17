// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

/*
✅ #1- Get and check the input
✅ #2- Create a byte buffer and use it as the output
✅ #3- Write input to the buffer as it is and print it
✅ #4- Detect the link
#5- Mask the link
#6- Stop masking when whitespace is detected
#7- Put a http:// prefix in front of the masked link
*/

package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {

	var mask rune = '*'
	httpMark := []byte("http://")

	// ✅ #1- Get and check the input
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Usage: Call program with text to check for spam url")
		return
	}

	// ✅ #2- Create a byte buffer and use it as the output
	textBuf := make([]byte, 0, len(args[0]))

	for i := 0; i < len(args[0]); i++ {
		textBuf = append(textBuf, args[0][i])
	}

	// ✅ #3- Write input to the buffer as it is and print it
	fmt.Printf("Type %[1]T, content = %[1]q, length = %[2]d, capacity = %[3]d\n", textBuf, len(textBuf), cap(textBuf))

	// ✅ #4- Detect the link
	fmt.Printf("%d\n", bytes.Index(textBuf, httpMark))
	var end int
	start := bytes.Index(textBuf, httpMark)
	if start > -1 {
		end = -1
		end = bytes.Index(textBuf[start+len(httpMark)-1:], []byte(" "))
		fmt.Printf("start+len(httpMark)-1 = %d, Link start = %d, link end = %d\n", start+len(httpMark)-1, start, end)
		if end == -1 {
			end = bytes.Index(textBuf, []byte("\t"))
			if end == -1 {
				end = bytes.Index(textBuf, []byte("\n"))
			}
		}
	}
	if start+len(httpMark)+end > start {
		for i := start; i < start+len(httpMark)+end-1; i++ {
			textBuf[i] = byte(mask)
		}
	}
	fmt.Printf("start+len(httpMark)-1 = %d, Link start = %d, link end = %d\n", start+len(httpMark)-1, start, end)
	fmt.Printf("Masked Output: Type %[1]T, content = %[1]q, length = %[2]d, capacity = %[3]d\n", textBuf, len(textBuf), cap(textBuf))
	fmt.Println(string(textBuf))
}
