package goherokuname

import (
	"fmt"
	"math/rand"
	"testing"
)

func init() {
	rand.Seed(99)
}

func ExampleHaikunate() {
	fmt.Println(Haikunate())
	// Output: dawn-pond-0223
}

func ExampleHaikunateHex() {
	fmt.Println(HaikunateHex())
	// Output: delicate-dawn-9869
}

func ExampleHaikunateCustom() {
	fmt.Println(HaikunateCustom("+", 5, "abcd"))
	// Output: tiny+limit+ddadc
}

func ExampleUbuntu() {
	out, _ := Ubuntu("-", "d")
	fmt.Println(out)
	// Output: dawn-dawn
}

func TestUbuntuInvalidLetter(t *testing.T) {
	_, err := Ubuntu("-", "1")
	if err == nil {
		t.Errorf("expected error missing")
	}
}
