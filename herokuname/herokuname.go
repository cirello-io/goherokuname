// Comand herokuname runs the goherokuname package in CLI.
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
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"cirello.io/goherokuname"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

var hex = flag.Bool(`hex`, false, `use hexadecimal values instead of decimals.`)
var digits = flag.Int(`digits`, 4, `how many digits after the names.`)
var separator = flag.String(`separator`, "-", `word separator.`)
var ubuntu = flag.String(`ubuntu`, "", `match initial letters (Ubuntu-style).`)
var listenHTTP = flag.String(`http`, "", `start HTTP server.`)

func main() {
	flag.Parse()

	if *listenHTTP != "" {
		http.HandleFunc("/about", serveAbout)
		http.HandleFunc("/ubuntu/", serveUbuntu)
		http.HandleFunc("/long", serveHerokuLong)
		http.HandleFunc("/short", serveHerokuShort)
		http.HandleFunc("/simple", serveHerokuSimple)
		http.HandleFunc("/", serveIndex)
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

func serveAbout(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://godoc.org/cirello.io/goherokuname", http.StatusSeeOther)
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

func serveHerokuSimple(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, goherokuname.HaikunateCustom("-", 0, ""))
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<html><body>
	<h1>Name generator as a service</h1>
	<ul>
		<li><a href="/about">about</li>
		<li><a href="/ubuntu/">ubuntu (e.g.: propellant-pomatomus)</li>
		<li><a href="/long">long (e.g.: overnice-minah-86758bf5)</li>
		<li><a href="/short">short (e.g.: cutaneal-muhammadanism-313e)</li>
		<li><a href="/simple">simple (e.g.: black-block)</li>
	</ul>
	</body></html>`)
}
