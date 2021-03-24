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
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Provide a directory")
		return
	}

	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	var names []byte

	for _, file := range files {
		if file.Size() == 0 {
			name := file.Name()

			fmt.Println("cap(names)", cap(names))
			names = append(names, name...)
			names = append(names, '\n')
		}
	}

	err = ioutil.WriteFile("out.txt", names, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", names)
}

// See: https://www.tutorialspoint.com/unix/unix-file-permission.htm
// See: http://permissions-calculator.org/
