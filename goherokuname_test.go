// Copyright 2019 github.com/ucirello and https://cirello.io. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to writing, software distributed
// under the License is distributed on a "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.

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
