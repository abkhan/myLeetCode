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

// digits are also in string
// return slice of strings that is combination of letters
func letterCombinations(digits string) []string {
	if len(digits) < 1 {
		return []string{}
	}

	// make array of digits
	var inputLetters [][]string
	for i := 0; i < len(digits); i++ {
		inputLetters = append(inputLetters, intLetterMap[digits[i]])
	}
	fmt.Printf("Array of required letters: %+v\n", inputLetters)

	return makes(inputLetters, 0)
}

// each letter is a string, input has slice of combination of letter
func makes(inl [][]string, ci int) []string {
	if ci == len(inl)-1 {
		return inl[len(inl)-1]
	}

	mids := makes(inl, ci+1) // all after current index
	these := inl[ci]         // current index

	// now multiply
	var retlist []string
	for _, th := range these { // for each letter of current index
		for _, mid := range mids { // for each string of combinations
			retlist = append(retlist, th+mid) // add the letter to each string
		}
	}

	return retlist
}
