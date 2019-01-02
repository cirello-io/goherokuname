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

package goherokuname

var (
	adjectivesIdx map[byte][]int
	nounsIdx      map[byte][]int
)

func init() {
	adjectivesIdx = extractIndexes(adjectives)
	nounsIdx = extractIndexes(nouns)
}

func extractIndexes(dict []string) map[byte][]int {
	idx := map[byte][]int{
		'a': {0},
	}

	lastLetter := byte('a')
	var i int
	var w string
	for i, w = range dict {
		if w[0] != lastLetter {
			idx[lastLetter] = append(idx[lastLetter], i-1)
			idx[w[0]] = []int{i}
			lastLetter = w[0]
		}
	}
	idx[lastLetter] = append(idx[lastLetter], i)
	return idx
}
