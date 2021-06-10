package main

import "log"

func main() {
	for _, n := range []int{3, 33, 41, 59, 61, 77, 143, 167, 371} {
		log.Printf("checkPrime for %d: %+v", n, checkPrime(n))
	}
}

func checkPrime(n int) bool {
	for x := 2; x*x <= n; x++ {
		if n%x == 0 {
			return false
		}
	}
	return true
}
