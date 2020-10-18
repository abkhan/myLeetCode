package main

import "fmt"

func main() {
	fmt.Printf("mySqrt:6:Res:%d\n", mySqrt(6))
	fmt.Printf("mySqrt:1267741:Res:%d\n", mySqrt(1267741))
}

func mySqrt(x int) int {

	if x == 0 {
		return 0
	}
	if x < 4 {
		return 1
	}

	for a := 2; a < 1000000; a++ {
		mul := a * a
		if mul == x {
			return a
		}
		if mul > x {
			return a - 1
		}
	}

	return -1
}
