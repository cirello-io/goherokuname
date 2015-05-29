// Package goherokuname generates Heroku-like random names. Such "block-black-0231".
//
// It supports both decimal and hexadecimal suffixes. It supports customized
// rendering, in which the delimiter, the length of suffix and the acceptable
// characters for suffix are tweakable.
package goherokuname

import (
	"math/rand"
	"time"
)

var seed int64

// HaikunateCustom generates a Heroku-like random name in which the delimiter,
// length and acceptable characters for suffix are tweakable.
func HaikunateCustom(delimiter string, toklen int, tokchars string) string {
	var lseed int64
	lseed = seed
	if 0 == lseed {
		lseed = time.Now().UnixNano()
	}
	r := rand.New(rand.NewSource(lseed))

	noun := nouns[r.Intn(len(nouns))]
	adjective := adjectives[r.Intn(len(adjectives))]

	token := ""
	for i := 0; i < toklen; i++ {
		token += string(tokchars[r.Intn(len(tokchars))])
	}

	return adjective + delimiter + noun + delimiter + token
}

// Haikunate generate standard Heroku-like random name, with "-" as delimiter,
// decimal with 4 digits.
func Haikunate() string {
	return HaikunateCustom("-", 4, "0123456789")
}

// HaikunateHex generate standard Heroku-like random name, with "-" as
// delimiter, hexadecimal with 4 digits.
func HaikunateHex() string {
	return HaikunateCustom("-", 4, "0123456789abcdef")
}
