package main

import "log"

func main() {
	tint := 0
	log.Printf("rev of %d, %d", tint, reverse(tint))
}

func reverse(x int) int {

	if x == 0 {
		return 0
	}

	if x > 2147483647 {
		return 0
	}

	neg := false
	if x < 0 {
		neg = true
		x = 0 - x
	}
	log.Printf("Input Value: %d", x)
	log.Printf("Input Value negative is %+v", neg)

	var iarr []int
	for i := 0; ; i++ {
		if x < 10 {
			iarr = append(iarr, x)
			log.Printf("End: Value at %d: %d", i, iarr[i])
			break
		}
		iarr = append(iarr, x%10)
		x = x / 10
		log.Printf("Value after %d: %d, >> x: %d", i, iarr[i], x)
	}
	log.Printf("Int Array: %+v", iarr)

	iarr = reverseInts(iarr)

	for i := len(iarr) - 1; i >= 0; i-- {
		tv := 1
		for j := 0; j < i; j++ {
			tv *= 10
		}
		iarr[i] *= tv
		log.Printf("Value at %d: %d", i, iarr[i])
	}

	retval := 0
	for i := 0; i < len(iarr); i++ {
		retval += iarr[i]
	}
	if retval > 2147483647 {
		return 0
	}
	if neg {
		retval = 0 - retval
	}
	return retval
}

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}
