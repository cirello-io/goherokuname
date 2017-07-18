package goherokuname

import "fmt"

// TODO: use go generate
func extractIndexes() {
	adjIdx := map[byte][]int{
		'a': {0},
	}
	lastAdjLetter := byte('a')
	var i int
	var adj string
	for i, adj = range adjectives {
		if adj[0] != lastAdjLetter {
			adjIdx[lastAdjLetter] = append(adjIdx[lastAdjLetter], i-1)
			adjIdx[adj[0]] = []int{i}
			lastAdjLetter = adj[0]
		}
	}
	adjIdx[lastAdjLetter] = append(adjIdx[lastAdjLetter], i)

	for letter, idxs := range adjIdx {
		fmt.Printf("'%v':{%v,%v},\n", string(letter), idxs[0], idxs[1])
	}

	fmt.Println("----")

	nounIdx := map[byte][]int{
		'a': {0},
	}
	lastNounLetter := byte('a')
	i = 0
	var noun string
	for i, noun = range nouns {
		if noun[0] != lastNounLetter {
			nounIdx[lastNounLetter] = append(nounIdx[lastNounLetter], i-1)
			nounIdx[noun[0]] = []int{i}
			lastNounLetter = noun[0]
		}
	}
	nounIdx[lastNounLetter] = append(nounIdx[lastNounLetter], i)

	for letter, idxs := range nounIdx {
		fmt.Printf("'%v':{%v,%v},\n", string(letter), idxs[0], idxs[1])
	}
}
