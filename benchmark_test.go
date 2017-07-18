package goherokuname_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"cirello.io/goherokuname"
)

func BenchmarkHaikunate(b *testing.B) {
	var resp string
	for i := 0; i < b.N; i++ {
		resp = goherokuname.Haikunate()
	}
	fmt.Fprintln(ioutil.Discard, resp)
}

func BenchmarkHaikunateHex(b *testing.B) {
	var resp string
	for i := 0; i < b.N; i++ {
		goherokuname.HaikunateHex()
	}
	fmt.Fprintln(ioutil.Discard, resp)
}

func BenchmarkHaikunateCustom(b *testing.B) {
	var resp string
	for i := 0; i < b.N; i++ {
		resp = goherokuname.HaikunateCustom("+", 5, "abcd")
	}
	fmt.Fprintln(ioutil.Discard, resp)
}

func BenchmarkUbuntu(b *testing.B) {
	var resp string
	var err error
	for i := 0; i < b.N; i++ {
		resp, err = goherokuname.Ubuntu("-", "a")
		if err != nil {
			b.Fatal(err)
		}
	}
	fmt.Fprintln(ioutil.Discard, resp)
}
