package main

import "log"

func main() {
	rs := longestPalindrome("c")

	log.Printf("longestPalindrome ret: %s", rs)
}

func longestPalindrome(s string) string {
	var allpal []string

	if len(s) < 2 {
		return s
	}

	for idx := 0; idx < len(s)-1; idx++ {
		for subidx := 1; subidx <= len(s)-idx; subidx++ {
			to := idx + subidx
			subs := s[idx:to]
			if isPalin(subs) {
				allpal = append(allpal, subs)
			}
		}
	}

	finalstr := ""
	for _, pal := range allpal {
		if len(pal) > len(finalstr) {
			finalstr = pal
		}
	}
	return finalstr
}

func isPalin(s string) bool {
	log.Printf("isPalin: Input> [%s]", s)
	odd := len(s) % 2
	halflen := len(s) / 2

	dni := halflen - 1
	upi := halflen + odd

	for ; dni >= 0; dni-- {
		if s[upi] != s[dni] {
			return false
		}
		upi++
	}
	return true
}
