package main

import (
	"fmt"
)

var testvals = [][]int{
	{10, 3, 3},
	{7, -2, -3},
	{0, 102, 0},
	{2345432, 3, 781810},
	{1619, -411, -3},
	{0, -5102, 0},
	{2, -2, -1},
	{2, 2, 1},
	{-20, -20, 1},
	{-2, 2, -1},
	{-2147483648, -1, 2147483647},
}

func main() {

	for ix, testval := range testvals {
		fmt.Printf("%d: Input>> %+v\n", ix, testval)
		at := divide(testval[0], testval[1])
		if at != testval[2] {
			fmt.Printf("    \t\tfailed>> %+v\n", at)
		}
		fmt.Printf("    Output>> %+v\n\n", at)
	}
	fmt.Println("Done")
}

func divide(dividend int, divisor int) int {

	const MaxUint = ^uint32(0)
	const MaxInt = int32(MaxUint >> 1)

	if dividend == 0 {
		return 0
	}

	dendneg, sorneg := false, false
	if dividend < 0 {
		dendneg = true
		dividend = 0 - dividend
	}
	if divisor < 0 {
		sorneg = true
		divisor = 0 - divisor
	}
	if dendneg == true && sorneg == true {
		dendneg = false
		sorneg = false
	}

	quot := 0

	//fmt.Printf("divid: %d, visor: %d .. MaxInt: %d", dividend, divisor, MaxInt)
	if divisor == 1 {
		quot = dividend
	} else if dividend == divisor {
		quot = 1
	} else {
		for dividend >= divisor {
			dividend -= divisor
			quot++
		}
	}

	if dendneg == true || sorneg == true {

		if quot > int(MaxInt)+1 {
			quot = int(MaxInt) + 1
		}

		quot = 0 - quot
	} else {
		if quot > int(MaxInt) {
			quot = int(MaxInt)
		}
	}

	return quot
}
