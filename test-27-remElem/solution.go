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
		len := removeElement(testval, 3)
		fmt.Printf("[%d] Output>> %+v\n\n", len, testval)
	}
	fmt.Println("Done")
}

func removeElement(nums []int, val int) int {
	if len(nums) < 1 {
		return 0
	}
	//if len(nums) < 2 {
	//	if nums[0] == val {
	//		return 0
	//	}
	//}
	var writeIndex = 0

	for readIndex := 0; readIndex < len(nums); readIndex++ {
		if nums[readIndex] == val {
			continue
		}
		nums[writeIndex] = nums[readIndex]
		writeIndex++
	}

	return writeIndex
}
