package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {

	var current, prev *ListNode

	current = head
	var firsttime bool = true
	for {

		if current == nil || current.Next == nil {
			break
		}

		a := current      // save current
		b := current.Next // rhs
		current = b.Next  // move current

		fmt.Printf("a: %d, b: %d\n", a.Val, b.Val)

		// switch
		if !firsttime {
			prev.Next = b
		}
		b.Next = a
		a.Next = current
		prev = a

		if firsttime {
			firsttime = false
			head = b
		}
	}

	return head
}
