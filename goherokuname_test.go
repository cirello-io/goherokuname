package goherokuname

import "fmt"

func ExampleHaikunate() {
	fmt.Println(Haikunate())
	// Output: waterfall-yellow-0312
}

func ExampleHaikunateHex() {
	fmt.Println(HaikunateHex())
	// Output: unit-red-0a12
}

func ExampleHaikunateCustom() {
	fmt.Println(HaikunateCustom("+", 5, "abcd"))
	// Output: term+twilight+abadc
}
