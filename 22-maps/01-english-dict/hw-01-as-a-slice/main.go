package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[english word] -> [german word]")
		return
	}
	query := args[0]

	english := []string{"good", "great", "perfect"}
	german := []string{"gut", "großartig", "perfekt"}

	// O(n) -> Inefficient: Depends on 'n'.
	for i, w := range english {
		if query == w {
			fmt.Printf("%q means %q\n", w, german[i])
			return
		}
	}

	fmt.Printf("%q not found\n", query)
}
