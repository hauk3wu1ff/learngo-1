package main

import (
	"fmt"
	"os"
)

func main() {
	// #2: Get the key from CLI
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[english Word] -> [german word]")
		return
	}
	query := args[0]

	// #1 Empty Map literal
	//dic := map[string]string{}

	// #3: Map literal
	dict := map[string]string{
		"good":    "nice",
		"great":   "großartig",
		"perfect": "perfekt",
	}

	dict["up"] = "hoch"     //add a new pair
	dict["down"] = "runter" // add a new pair
	dict["good"] = "gut"    // #5: overwrites the value at key: "good"
	dict["mistake"] = ""    // #8: a key with a zero-value

	// 10#: comma ok in a short if
	if value, ok := dict[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}
	fmt.Printf("%q not found.\n", query)

	fmt.Printf("# of Keys: %d\n", len(dict))

	// #13: compare a map using its printed output
	copied := map[string]string{"up": "hoch", "down": "runter",
		"mistake": "", "good": "gut", "great": "großartig",
		"perfect": "perfekt"}

	first := fmt.Sprintf("%s", dict)
	second := fmt.Sprintf("%s", copied)

	if first == second {
		fmt.Println("Maps are equal")
	}

	// #12: printing a map (ordered output since Go 1.12)
	fmt.Printf("%#v\n", dict)

	// #11
	for k, v := range dict {
		fmt.Printf("%q means %#v\n", k, v)
	}

	// #9: Check for non-existebt key: with comma, ok
	_, ok := dict[query]
	if !ok {
		fmt.Printf("%q not found.\n", query)
	}

	// #6: getting values from a map using keys directly
	fmt.Println("good           ->", dict["good"])
	fmt.Println("great           ->", dict["great"])
	fmt.Println("perfect           ->", dict["perfect"])

	// #2b: retrieve values by key - O(1) efficency
	value := dict[query]
	fmt.Printf("%q means %#v\n", query, value)

}
