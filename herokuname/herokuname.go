//
// Usage of herokuname:
//   -digits int
//        	how many digits after the names. (default 4)
//   -hex
//        	use hexadecimal values instead of decimals.
//   -http string
//        	start HTTP server.
//   -separator string
//        	word separator. (default "-")
//   -ubuntu string
//        	match initial letters (Ubuntu-style).
//
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"cirello.io/goherokuname"
)

var hex = flag.Bool(`hex`, false, `use hexadecimal values instead of decimals.`)
var digits = flag.Int(`digits`, 4, `how many digits after the names.`)
var separator = flag.String(`separator`, "-", `word separator.`)
var ubuntu = flag.String(`ubuntu`, "", `match initial letters (Ubuntu-style).`)
var listenHTTP = flag.String(`http`, "", `start HTTP server.`)

func main() {
	flag.Parse()

	if *listenHTTP != "" {
		http.HandleFunc("/ubuntu/", serveUbuntu)
		http.HandleFunc("/long", serveHerokuLong)
		http.HandleFunc("/", serveHerokuShort)
		if err := http.ListenAndServe(*listenHTTP, nil); err != nil {
			log.SetPrefix("")
			log.SetFlags(0)
			log.Fatal(err)
		}
		os.Exit(0)
	}

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

func serveUbuntu(w http.ResponseWriter, r *http.Request) {
	letter := strings.TrimPrefix(r.URL.String(), "/ubuntu/")
	name, err := goherokuname.Ubuntu("-", letter)
	if err != nil {
		http.Error(w, fmt.Sprintln(http.StatusText(http.StatusBadRequest), err), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, name)
}

func serveHerokuLong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, goherokuname.HaikunateCustom("-", 8, "0123456789abcdef"))
}

func serveHerokuShort(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, goherokuname.HaikunateHex())
}
