package goherokuname

import "fmt"

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
