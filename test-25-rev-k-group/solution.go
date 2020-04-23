package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

var lval = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

func main() {
	var head, curr *ListNode
	for ix := 0; ix < len(lval); ix++ {
		aNp := &ListNode{Val: lval[ix]}
		if curr == nil {
			curr = aNp
			head = curr
		} else {
			curr.Next = aNp
			curr = curr.Next
		}
	}
	printLL(head)
	head = reverseKGroup(head, 3)
	fmt.Print("after ...\n")
	printLL(head)
}

func printLL(head *ListNode) {
	fmt.Print("[ ")
	for head != nil {
		fmt.Printf("%d, ", head.Val)
		if head.Next == nil {
			break
		}
		head = head.Next
	}
	fmt.Print("]\n")
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	if k < 2 {
		return head
	}

	var resSlice []*ListNode

	currS := head
	next := head
	lastIter := false

	for {

		// take k and call reverse
		for ix := 0; ix < k; ix++ {
			if next == nil {
				lastIter = true
				break
			}
			next = next.Next
		}

		//fmt.Print(">>>")
		//printLL(currS)
		rev := reverseElem(currS, k)
		//fmt.Print("<<<\n")

		// Take the return and append to resultSlice
		for ix := 0; ix < k; ix++ {
			if rev == nil {
				lastIter = true
				break
			}
			resSlice = append(resSlice, rev)
			rev = rev.Next
		}

		// move current up
		currS = next

		if lastIter {
			break
		}
	}

	// make a new LL
	if len(resSlice) > 0 {
		head = resSlice[0]
		currS = head
		for ix := 1; ix < len(resSlice); ix++ {
			currS.Next = resSlice[ix]
			currS = currS.Next
		}
		currS.Next = nil
	}

	return head
}

func reverseElem(head *ListNode, k int) *ListNode {

	if k < 2 {
		return head
	}
	// lets make sure we have all
	start := head
	count := 0
	for ix := 0; ix < k; ix++ {
		if start != nil {
			start = start.Next
			count++
		}
	}
	if count < k {
		//fmt.Print("not enough, ret same")
		return head
	}

	start = head
	var kg []*ListNode
	for ix := 0; ix < k; ix++ {
		if start != nil {
			kg = append(kg, start)
			start = start.Next
		} else {
			if ix < k-1 {
				return head
			}
		}
	}

	head = kg[len(kg)-1]
	start = head
	for ix := len(kg) - 2; ix >= 0; ix-- {
		start.Next = kg[ix]
		start = start.Next
	}
	return head
}
