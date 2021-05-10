package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Printf("[english word] -> [german word]")
		return
	}

	query := args[0]
	dict := map[string]string{
		"good":    "gut",
		"great":   "großartig",
		"perfect": "perfekt",
		"awesome": "genial", // #5
	}

	german := make(map[string]string, len(dict))
	for k, v := range dict {
		german[v] = k
	}

	delete(dict, "awesome")
	delete(dict, "awesone") // no-op
	delete(dict, "nonexisting")
	fmt.Printf("english: %q\ngerman: %q\n", dict, german)
	fmt.Printf("%q means %#v\n", query, dict[query])

	if value, ok := dict[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}

}
