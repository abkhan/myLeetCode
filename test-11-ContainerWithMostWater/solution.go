package main

import "log"

var tstr string

func main() {
	testLines := []int{0, 7}
	log.Printf("For lines %+v, maxArea is %d\n", testLines, maxArea(testLines))
	log.Print("-----\n")
	testLines = []int{1, 8, 6, 2, 0, 4, 8, 3, 7}
	log.Printf("For lines %+v, maxArea is %d\n", testLines, maxArea(testLines))
	log.Print("-----\n")
}

func maxArea(height []int) int {

	if len(height) < 2 {
		return 0
	}
	if len(height) == 2 {
		return minOf(height[0], height[1])
	}

	currMax := 0
	for tox := len(height) - 1; tox > 0; tox-- {
		for idx := 0; idx < tox; idx++ {
			dist := tox - idx
			effh := minOf(height[tox], height[idx])
			area := dist * effh
			if area > currMax {
				currMax = area
			}
		}
	}
	return currMax
}

func minOf(a, b int) int {
	if a < b {
		return a
	}
	return b
}
