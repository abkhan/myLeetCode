package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func onePass_removeNthFromEnd(head *ListNode, n int) *ListNode {
	var lna []*ListNode // use a list of all nodes to skip
	nxt := head
	for {
		if nxt != nil {
			lna = append(lna, nxt) // go over linked list and create a list of all nodes
			nxt = nxt.Next
		} else {
			break
		}
	}
	skip := len(lna) - n
	if skip == 0 {
		return head.Next
	}

	lna[skip-1].Next = lna[skip].Next // just update in the list-of-all-nodes and lined list will be changed
	return head
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if n < 1 {
		return head
	}
	if head == nil {
		return head
	}

	lsize := 1
	var nxt *ListNode = head.Next
	// first find out the size
	for {
		if nxt != nil {
			lsize++
			nxt = nxt.Next
		} else {
			break
		}
	}

	// now skip size-n nodes
	skip := lsize - n
	if skip == 0 {
		return head.Next
	}

	nxt = head
	for skip > 1 {
		nxt = nxt.Next
		skip--
	}
	toSkip := nxt.Next
	nxt.Next = toSkip.Next

	return head
}
