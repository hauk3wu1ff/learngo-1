package main

import "fmt"

const nihongo = "日本語"

func main() {
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
}
