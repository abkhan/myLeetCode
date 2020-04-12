package main

import (
	"fmt"
	"log"
)

var testvals = [][]string{
	{"flower", "flow", "flight"},
	{"whatevermaybe", "whatever"},
	{"dog", "racecar", "car"},
	{"dog", "dogger", "doggest"},
	{"hi", "racecar", ""},
	{},
	{"whatever"},
}

func main() {

	for _, tvals := range testvals {
		log.Printf("For input %+v, longest common profix is [%s]\n", tvals, longestCommonPrefix(tvals))
		log.Print("-----\n")
	}
}

func longestCommonPrefix(strs []string) string {

	rets := ""
	minStrLen := 999

	if len(strs) < 1 {
		return ""
	}
	if len(strs) < 2 {
		return strs[0]
	}

	for _, s := range strs {
		if len(s) < minStrLen {
			minStrLen = len(s)
		}
	}

	fmt.Printf("minStrLen: %d\n", minStrLen)

	for i := 0; i < minStrLen; i++ {
		match := true
		ch := strs[0][i]
		for _, s := range strs {
			if ch != s[i] {
				match = false
				break
			}
		}
		if !match {
			break
		}
		rets = rets + string(ch)
	}
	return rets
}
