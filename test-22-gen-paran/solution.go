package main

import (
	"log"
)

var testvals = []int{
	-1, 0, 1, 2, 3, 4, 5, 6,
}

func main() {

	for _, tvals := range testvals {
		log.Printf("for %d paran [%+v]\n", tvals, generateParenthesis(tvals))
		log.Print("-----\n")
	}
}

func generateParenthesis(n int) []string {

	if n < 1 {
		return []string{}
	}
	if n == 1 {
		return []string{"()"}
	}

	allP := createP(n * 2)
	return validateP(allP)
}

func validateP(strs []string) []string {
	var ts, rv []string
	for _, ss := range strs {
		if ss[0] != ')' {
			ts = append(ts, ss)
		}
	}

	for _, s := range ts {
		if validateS(s) {
			rv = append(rv, s)
		}
	}
	return rv
}

// true means OK
func validateS(s string) bool {

	mustStart := true
	endPending := 0
	for x := 0; x < len(s); x++ {
		if s[x] == '(' {
			endPending++
			mustStart = false
		} else {
			if mustStart == true {
				return false
			}
			endPending--
			if endPending == 0 {
				mustStart = true
			}
		}
	}
	if endPending != 0 {
		return false
	}
	return true
}

func createP(n int) []string {
	if n == 1 {
		return []string{")"}
	}

	var retval []string
	strs := createP(n - 1)
	for _, v := range []string{"(", ")"} {
		for _, s := range strs {
			retval = append(retval, v+s)
		}
	}
	return retval
}
