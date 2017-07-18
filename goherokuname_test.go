package goherokuname

import (
	"fmt"
	"testing"
)

func init() {
	seed = 99
}

func ExampleHaikunate() {
	fmt.Println(Haikunate())
	// Output: dawn-pond-0223
}

func ExampleHaikunateHex() {
	fmt.Println(HaikunateHex())
	// Output: dawn-pond-2a6b
}

func ExampleHaikunateCustom() {
	fmt.Println(HaikunateCustom("+", 5, "abcd"))
	// Output: dawn+pond+cccdb
}

func ExampleUbuntu() {
	out, _ := Ubuntu("-", "d")
	fmt.Println(out)
	// Output: damp-dawn
}

func TestUbuntuInvalidLetter(t *testing.T) {
	_, err := Ubuntu("-", "1")
	if err == nil {
		t.Errorf("expected error missing")
	}
}
