package main

import "log"

var times int

func main() {
	for _, n := range []int{3, 33, 41, 59, 61, 77, 143, 167, 371} {
		log.Printf("count factorial for %d: %d", n, countFactorial(n))
		log.Printf("   >>> count %d", times)
		times = 0
	}

}

func countFactorial(n int) int64 {
	times++
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 1
	}
	return int64(n) * countFactorial(n-1)
}
