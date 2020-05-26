package main

import (
	"fmt"
	"strings"
)

var testvals = [][]string{
	{"needle", "dl"},
	{"hi", ""},
	{"what", "no"},
	{"", ""},
}

func main() {

	for ix, testval := range testvals {
		fmt.Printf("%d: Input>> %+v\n", ix, testval)
		at := strStr(testval[0], testval[1])
		fmt.Printf("    Output>> %+v\n\n", at)
	}
	fmt.Println("Done")
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	return strings.Index(haystack, needle)
}
