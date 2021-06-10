package main

import "log"

var memo map[int]int

func main() {
	memo = make(map[int]int)

	for _, n := range []int{3, 7, 12, 101} {
		log.Printf("allfibs for %d: %+v", n, allfibs(n))
	}
}

func allfibs(n int) []int {
	retval := []int{}
	for i := 0; i <= n; i++ {
		retval = append(retval, fib(i))
	}
	return retval
}

func fib(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if savedvalue, found := memo[n]; found {
		return savedvalue
	}

	memo[n] = fib(n-1) + fib(n-2)
	return memo[n]
}
