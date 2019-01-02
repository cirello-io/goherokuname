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
