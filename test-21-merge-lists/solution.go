package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var rv, cp *ListNode // retrun value, and copy pointer

	for {
		if l1 == nil || l2 == nil {
			if l1 == nil {
				if cp == nil {
					cp = l2
					rv = cp
				} else {
					cp.Next = l2
				}
				break
			}
			if l2 == nil {
				if cp == nil {
					cp = l1
					rv = cp
				} else {
					cp.Next = l1
				}
				break
			}
		}

		if l1.Val < l2.Val {
			on := l1
			l1 = l1.Next
			on.Next = nil
			if cp == nil {
				cp = on
				rv = cp
			} else {
				cp.Next = on
				cp = cp.Next
			}
		} else {
			on := l2
			l2 = l2.Next
			on.Next = nil
			if cp == nil {
				cp = on
				rv = cp
			} else {
				cp.Next = on
				cp = cp.Next
			}
		}
	}

	return rv
}
