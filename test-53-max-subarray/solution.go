package main

import "fmt"

func main() {
	fmt.Printf("maxSubArray: %d", maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	nlen := len(nums)
	maxIndex := 0
	maxCount := nlen
	maxValue := getSum(nums, maxIndex, maxCount)

	for ltr := 0; ltr < nlen; ltr++ {
		ltrsum := nums[ltr]
		if ltrsum > maxValue {
			maxIndex = ltr
			maxCount = nlen - ltr
			maxValue = ltrsum
		}

		for sub := ltr + 1; sub < nlen; sub++ {
			ltrsum += nums[sub]
			if ltrsum > maxValue {
				maxIndex = ltr
				maxCount = sub - ltr + 1
				maxValue = ltrsum
			}
		}
	}

	return maxValue
}

func getSum(nums []int, from int, count int) int {
	sum := nums[from]
	for x := 1; x < count; x++ {
		sum += nums[from+x]
	}
	fmt.Printf("From: %d, count: %d, Sum: %d\n", from, count, sum)
	return sum
}
