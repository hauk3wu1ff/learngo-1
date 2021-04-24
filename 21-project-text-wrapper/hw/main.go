package main

import (
	"fmt"
	"unicode"
)

func main() {
	const text = `Chet Baker, mit bürgerlichem Namen Chesney Henry Baker Jr. (* 23. Dezember 1929 in Yale, Oklahoma; † 13. Mai 1988 in Amsterdam), war ein US-amerikanischer Jazzmusiker (Trompete, Flügelhorn), Sänger und Komponist.

Im Alter von zehn Jahren erhielt Baker von seinem Vater, der Gitarrist war, eine Posaune.[1][2] Der junge Chet tauschte sie gegen eine Trompete ein. Innerhalb kürzester Zeit konnte er einfache Swing-Melodien nachspielen und lernte nebenbei als Autodidakt das Improvisieren. Seine besondere Fähigkeit bestand darin, harmonische Zusammenhänge schnell zu erfassen und harmonische Vorgaben melodisch zu umspielen. Baker hat in seinem Leben nur wenig komponiert.

Zunächst spielte er im Schulorchester, sodann bei öffentlichen Tanzveranstaltungen. 1946 meldete er sich zum Militär und kam zur 298th Army Band nach Berlin. 1948 wurde er ausgemustert. Er studierte Theorie und Harmonie im El Camino College in Los Angeles, abends spielte er in Clubs.
 `

	const maxWidth = 40

	var lw int // line width

	for _, r := range text {
		fmt.Printf("%c", r)

		// Printing a rune increases the line length
		// That is why we use lw++ to increase lw by 1, before
		// checking lw in the case condition
		switch lw++; {
		case lw > maxWidth && r != '\n' && unicode.IsSpace(r):
			// Print line-break when lw exceeds maxWidth
			fmt.Println()
			fallthrough
		case r == '\n':
			// We have to also reset lw after a '\n' in
			// the original text.
			lw = 0
		}
	}
}
