package main

import (
	"fmt"
	"log"
)

var testvals = [][]int{
	{1, 2, 3, 4, -2, -5, -1, 0, 4, -6, 7},
	{-2, -5, -1, 0, 4, -6, 7},
	{-6, 7},
	{7, -2, 3, 4, -2, -5, -1, 0, 4, -6, 7, 8},
	{7, 2, -9, -10, -11, 11, 0},
	{7, -2, 3, 4, -2, -5, -1, 0, 4, 7},
	{0, 0, 0, 0},
	{-15, 13, 6, -11, -4, 5, -13, 5, 3, 2, 6, -1, 4, 12, -10, -13, -7, -4, -5, 6, 9, -14, 1, -6, 13, 7, -8, 10, -4, 11, -8, -3, 1, 5, -7, 4, -13, -13, -5, -3, 4, -14, 11, -14, 5, -13, -12, 13, -10, -10, -4, -15, 13, 13, -14, 11, -3, -15, 6, 1, 3, 5, 13, -11, -5, -9, 1, -2, -14, 11, 10, 5, 4, -1, 6, -6, -7, 9, -15, -2, 7, -12, -10, 5, -14, 13, -6, -9, 6, 7, 7, -6, -2, -3, -9, 0, -5, 7, 5, -4, -5, -7, -13, 14, 7, 8, -15, 7, -5, -15, -10, 9},
}

func main() {

	for _, tvals := range testvals {
		log.Printf("For input %+v, 3Sum [%+v]\n", tvals, threeSum(tvals))
		log.Print("-----\n")
	}
}

func threeSum(nums []int) [][]int {
	var perms [][]int
	var retv [][]int

	if len(nums) < 3 {
		return retv
	}

	// first create a permutation
	for fi := 0; fi < len(nums)-2; fi++ {

		for se := fi + 1; se < len(nums)-1; se++ {

			for th := se + 1; th < len(nums); th++ {
				var onep []int = []int{nums[fi], nums[se], nums[th]}
				perms = append(perms, onep)
			}
		}
	}

	fmt.Printf("perms: %d\n", len(perms))

	// now create
	for _, onent := range perms {
		if onent[0]+onent[1]+onent[2] == 0 {

			onent = sort3sum(onent)
			// check if already there
			if !alreadythere(onent, retv) {
				retv = append(retv, onent)
			}
		}
	}

	fmt.Printf("retv: %d\n", len(retv))

	return retv
}

func alreadythere(e []int, l [][]int) bool {
	for _, each := range l {
		if each[0] == e[0] &&
			each[1] == e[1] &&
			each[2] == e[2] {
			return true
		}
	}
	return false
}

func sort3sum(e []int) []int {

	if e[0] > e[1] {
		t := e[1]
		e[1] = e[0]
		e[0] = t
	}

	if e[1] > e[2] {
		t := e[2]
		e[2] = e[1]
		e[1] = t
	}

	if e[0] > e[1] {
		t := e[1]
		e[1] = e[0]
		e[0] = t
	}

	return e
}
