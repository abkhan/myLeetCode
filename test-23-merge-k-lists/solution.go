package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {

	var rethead, curt *ListNode

	if len(lists) < 1 {
		return nil
	}
	if len(lists) < 2 {
		return lists[0]
	}
	for {

		var smallest int
		lowIx := -1
		someLeft := false
		for ix := 0; ix < len(lists); ix++ {
			if lists[ix] == nil {
				continue
			}
			someLeft = true

			if lowIx == -1 {
				smallest = lists[ix].Val
				lowIx = ix
				continue
			}

			if lists[ix].Val < smallest {
				smallest = lists[ix].Val
				lowIx = ix
				continue
			}
		}

		if someLeft == false {
			break
		}

		if rethead == nil {
			rethead = lists[lowIx]
		}
		if curt == nil {
			curt = rethead
		} else {
			curt.Next = lists[lowIx]
			curt = curt.Next
		}

		lists[lowIx] = lists[lowIx].Next
	}

	return rethead
}
