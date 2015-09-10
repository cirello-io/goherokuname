package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dericofilho/goherokuname"
)

var hex = flag.Bool(`hex`, false, `use hexadecimal values instead of decimals.`)
var digits = flag.Int(`digits`, 4, `how many digits after the names.`)
var separator = flag.String(`separator`, "-", `word separator.`)

func main() {
	flag.Parse()

	tokchars := "0123456789"
	if *hex {
		tokchars = tokchars + "abcdef"

	}

	fmt.Println(goherokuname.HaikunateCustom(*separator, *digits, tokchars))
	os.Exit(0)
}
