package main

import "log"

func main() {
	ti := 0
	log.Printf("----->[%d], palind: %+v", ti, isPalindrome(ti))
	ti = -919
	log.Printf("----->[%d], palind: %+v", ti, isPalindrome(ti))
	ti = 1
	log.Printf("----->[%d], palind: %+v", ti, isPalindrome(ti))
	ti = 1121
	log.Printf("----->[%d], palind: %+v", ti, isPalindrome(ti))
	ti = 11
	log.Printf("----->[%d], palind: %+v", ti, isPalindrome(ti))
	ti = 121
	log.Printf("----->[%d], palind: %+v", ti, isPalindrome(ti))
	ti = 971179
	log.Printf("----->[%d], palind: %+v", ti, isPalindrome(ti))
	ti = -121
	log.Printf("----->[%d], palind: %+v", ti, isPalindrome(ti))
	ti = 3391933
	log.Printf("----->[%d], palind: %+v", ti, isPalindrome(ti))
}

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	var iarr []int
	for i := 0; ; i++ {
		if x < 10 {
			iarr = append(iarr, x)
			break
		}
		iarr = append(iarr, x%10)
		x = x / 10
	}

	restr := ""
	for i := 0; i < len(iarr); i++ {
		restr += string(iarr[i] + '0')
	}
	return isPalin(restr)
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
