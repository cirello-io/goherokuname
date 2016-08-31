//
// Usage of herokuname:
//   -digits int
//        	how many digits after the names. (default 4)
//   -hex
//        	use hexadecimal values instead of decimals.
//   -separator string
//        	word separator. (default "-")
//   -ubuntu string
//        	match initial letters (Ubuntu-style).
//
package main

import (
	"flag"
	"fmt"
	"os"

	"cirello.io/goherokuname"
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
		name, err := goherokuname.Ubuntu(*separator, *ubuntu)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
		fmt.Println(name)
	} else {
		fmt.Println(goherokuname.HaikunateCustom(*separator, *digits, tokchars))
	}

	os.Exit(0)
}
