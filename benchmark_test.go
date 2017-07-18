package goherokuname_test

import (
	"testing"

	"cirello.io/goherokuname"
)

func BenchmarkHaikunateDeterministic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goherokuname.Haikunate()
	}
}

func BenchmarkHaikunateHexDeterministic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goherokuname.HaikunateHex()
	}
}

func BenchmarkHaikunateCustomDeterministic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goherokuname.HaikunateCustom("+", 5, "abcd")
	}
}

func BenchmarkHaikunateRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goherokuname.Haikunate()
	}
}

func BenchmarkHaikunateHexRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goherokuname.HaikunateHex()
	}
}

func BenchmarkHaikunateCustomRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goherokuname.HaikunateCustom("+", 5, "abcd")
	}
}
