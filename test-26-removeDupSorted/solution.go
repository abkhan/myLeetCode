package main

import "fmt"

var testvals = [][]int{
	{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
	{},
	{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
	{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4},
	{1, 1, 2},
	{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
	{-1},
	{1, 2},
}

func main() {

	for ix, testval := range testvals {
		fmt.Printf("%d: Input>> %+v\n", ix, testval)
		len := removeDuplicates(testval)
		fmt.Printf("[%d] Output>> %+v\n\n", len, testval)
	}
	fmt.Println("Done")
}

func removeDuplicates(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	if len(nums) < 2 {
		return 1
	}

	var prevValue = nums[0]
	var writeIndex = 1

	for readIndex := 1; readIndex < len(nums); readIndex++ {
		currentValue := nums[readIndex]
		if currentValue == prevValue {
			// case of a duplicate
			continue
		}
		nums[writeIndex] = currentValue
		prevValue = currentValue
		writeIndex++
	}

	return writeIndex
}
