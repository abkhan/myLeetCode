package main

import "fmt"

func main() {
	fmt.Printf("test1: [1,3,5,6], 5 >> %d\n", searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Printf("test1: [1,3,5,6], 2 >> %d\n", searchInsert([]int{1, 3, 5, 6}, 2))
}

func searchInsert(nums []int, target int) int {
	if len(nums) < 1 {
		return 0
	}
	for x := 0; x < len(nums); x++ {
		if nums[x] >= target {
			return x
		}
	}
	return len(nums)
}
