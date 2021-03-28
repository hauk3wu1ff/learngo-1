package main

import "fmt"

func main() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println(sample)

	// examine bytes using a loop
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n")

	// shorter way using string and %x printf verb
	fmt.Printf("%x\n", sample)

	// With space-flag in printf format
	fmt.Printf("% x\n", sample)

	// With quoted %q printf verb
	fmt.Printf("%q\n", sample)

	//If we are unfamiliar or confused by strange values in the string,
	// we can use the "plus" flag to the %q verb. This flag causes the
	// output to escape not only non-printable sequences, but also any
	// non-ASCII bytes, all while interpreting UTF-8. The result is that
	// it exposes the Unicode values of properly formatted UTF-8 that
	// represents non-ASCII data in the string:
	fmt.Printf("%+q\n", sample)
}
