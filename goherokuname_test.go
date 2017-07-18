package goherokuname_test

import (
	"fmt"
	"math/rand"
	"testing"

	"cirello.io/goherokuname"
)

func init() {
	rand.Seed(99)
}

func ExampleHaikunate() {
	fmt.Println(goherokuname.Haikunate())
	// Output: dawn-pond-0223
}

func ExampleHaikunateHex() {
	fmt.Println(goherokuname.HaikunateHex())
	// Output: delicate-dawn-9869
}

func ExampleHaikunateCustom() {
	fmt.Println(goherokuname.HaikunateCustom("+", 5, "abcd"))
	// Output: tiny+limit+ddadc
}

func ExampleUbuntu() {
	out, _ := goherokuname.Ubuntu("-", "d")
	fmt.Println(out)
	// Output: dawn-dawn
}

func TestUbuntuInvalidLetter(t *testing.T) {
	_, err := goherokuname.Ubuntu("-", "1")
	if err == nil {
		t.Errorf("expected error missing")
	}
}
