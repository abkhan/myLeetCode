package main

import (
	"fmt"
	"math"
	"sort"
)

type atest struct {
	input  []int
	result []int
}

var testvals = []atest{
	{[]int{1, 2, 3}, []int{1, 3, 2}},
	{[]int{3, 2, 1}, []int{1, 2, 3}},
	{[]int{1, 1, 5}, []int{1, 5, 1}},
	{[]int{4, 1, 2, 3}, []int{4, 1, 3, 2}},
	{[]int{4, 3, 7, 3}, []int{4, 7, 3, 3}},
	{[]int{1, 2, 3, 4}, []int{1, 2, 4, 3}},
	{[]int{2, 3, 4, 1}, []int{2, 4, 3, 1}},
	{[]int{1, 3, 2}, []int{2, 1, 3}},
}

func main() {

	for ix, testval := range testvals {
		fmt.Printf("%d: Input>> %+v\n", ix, testval)
		nextPermutation(testval.input)
		fmt.Printf("    Retval>> %+v,    Expected>> %+v\n\n", testval.input, testval.result)
	}
	fmt.Println("Done")
}

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	if isSame(nums) || isRevSorted(nums) {
		sort.Ints(nums)
		return
	}

	inputval := decValaue(nums)
	//fmt.Printf("Input Value: %d", inputval)

	// first pass, compare and change with next in line
	minR := nums[len(nums)-1]
	for ix := len(nums) - 1; ix > 0; ix-- {
		if nums[ix] < minR {
			minR = nums[ix]
		}
		if nums[ix] > nums[ix-1] && nums[ix-1] <= minR { //
			nums[ix-1], nums[ix] = nums[ix], nums[ix-1]
			fmt.Printf(" >> First Pass solution >>> Input Value: %d, newVal: %d\n", inputval, decValaue(nums))
			return
		}
	}

doneLoop:
	for ix := len(nums) - 1; ix > 1; ix-- {
		for iy := len(nums) - 2; iy > 0; iy-- {
			if nums[ix] > nums[iy] {
				numx := nums[ix] // save x val
				// move iy to right
				fmt.Printf(" >> Move: %d, to: %d\n", ix, iy)

				for im := len(nums) - 2; im >= iy; im-- {
					nums[im+1] = nums[im]
				}
				nums[iy] = numx

				fmt.Printf(" >> Input Value: %d, newVal: %d\n", inputval, decValaue(nums))
				break doneLoop
			}

		}

	} // end of doneLoop:

}

func isSame(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	f := nums[0]
	for ix := 1; ix < len(nums); ix++ {
		if nums[ix] != f {
			return false
		}
	}
	return true
}

func isRevSorted(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	c := nums[0]
	for ix := 1; ix < len(nums); ix++ {
		if nums[ix] > c { // if greater, not reverse sorted
			return false
		}
		c = nums[ix]
	}
	return true
}

func decValaue(il []int) int {
	ill := len(il)
	res := 0
	for ix := ill - 1; ix >= 0; ix-- {
		res += int(math.Pow10(ix)) * il[ill-ix-1]
	}
	return res
}
