//
// Usage of herokuname:
//   -digits int
//     	how many digits after the names. (default 4)
//   -hex
//     	use hexadecimal values instead of decimals.
//   -separator string
//     	word separator. (default "-")
//
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ccirello/goherokuname"
)

var hex = flag.Bool(`hex`, false, `use hexadecimal values instead of decimals.`)
var digits = flag.Int(`digits`, 4, `how many digits after the names.`)
var separator = flag.String(`separator`, "-", `word separator.`)
var ubuntu = flag.String(`ubuntu`, "", `match initial letters (Ubuntu-style).`)

func main() {
	flag.Parse()

	tokchars := "0123456789"
	if *hex {
		tokchars = tokchars + "abcdef"
	}

	if *ubuntu != "" {
		fmt.Println(goherokuname.Ubuntu(*separator, *ubuntu))
	} else {
		fmt.Println(goherokuname.HaikunateCustom(*separator, *digits, tokchars))
	}

	os.Exit(0)
}
