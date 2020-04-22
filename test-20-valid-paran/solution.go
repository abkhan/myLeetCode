package main

import (
	"log"
)

var testvals = []string{
	"[]", "[(})]", "()", "()[]{}", "(]", "{[]}", "([)]", "", "((((((({{[[(([]))]]}}))))}[)",
}

func main() {

	for _, tvals := range testvals {
		log.Printf("is %s valid [%+v]\n", tvals, isValid(tvals))
		log.Print("-----\n")
	}
}

func isValid(s string) bool {

	if len(s) < 1 {
		log.Printf("return for small size")
		return true
	}

	if len(s)%2 == 1 {
		log.Printf("return with wrong size")
		return false
	}

	var thestack []byte = []byte{s[0]}

	for ix := 1; ix < len(s); ix++ {
		cp := s[ix]
		if cp == '(' || cp == '[' || cp == '{' {
			thestack = append(thestack, cp)
			continue
		}

		if cp == ')' {
			if thestack[len(thestack)-1] == '(' {
				thestack = thestack[:len(thestack)-1]
			} else {
				return false
			}
			continue
		}
		if cp == ']' {
			if thestack[len(thestack)-1] == '[' {
				thestack = thestack[:len(thestack)-1]
			} else {
				return false
			}
			continue
		}
		if cp == '}' {
			if thestack[len(thestack)-1] == '{' {
				thestack = thestack[:len(thestack)-1]
			} else {
				return false
			}
			continue
		}
	}

	if len(thestack) > 0 {
		log.Printf("Stack @ the end: %+v", thestack)
		return false
	}
	return true
}
