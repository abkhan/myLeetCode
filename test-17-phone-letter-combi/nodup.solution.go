package main

import (
	"fmt"
	"log"
)

var testvals = []string{
	"23", "", "2", "99999", "789", "445",
}

var intLetterMap map[byte][]string = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func main() {

	for _, tvals := range testvals {
		log.Printf("For input [%s] combination is [%+v]\n", tvals, letterCombinations(tvals))
		log.Print("-----\n")
	}
}

func letterCombinations(digits string) []string {
	if len(digits) < 1 {
		return []string{}
	}

	// remove duplicates
	// by using a letter-map so that if there is a duplicate, it would not be a duplicate in output
	nodup := make(map[byte]bool)
	for i := 0; i < len(digits); i++ {
		nodup[digits[i]] = true
	}

	// make array of nodup
	var inputLetters [][]string
	for k := range nodup {
		inputLetters = append(inputLetters, intLetterMap[k])
	}
	fmt.Printf("Array of required letters: %+v\n", inputLetters)

	return makes(inputLetters, 0)
}

func makes(inl [][]string, ci int) []string {
	if ci == len(inl)-1 {
		return inl[len(inl)-1]
	}

	mids := makes(inl, ci+1)
	these := inl[ci]

	// now multiply
	var retlist []string
	for _, th := range these {
		for _, mid := range mids {
			retlist = append(retlist, th+mid)
		}
	}

	return retlist
}
