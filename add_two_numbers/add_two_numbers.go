package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(v1 *ListNode, v2 *ListNode) *ListNode {
	var first ListNode
	curr := &first
	var prev *ListNode
	rip := 0
	var l1, l2 = v1, v2
	var l1v, l2v int
	for l1 != nil || l2 != nil {
		// Preparo gli addendi
		if l1 != nil && l2 != nil {
			l1v, l2v = l1.Val, l2.Val
			l1, l2 = l1.Next, l2.Next
		} else if l1 != nil {
			l1v, l2v = l1.Val, 0
			l1 = l1.Next
		} else if l2 != nil {
			l1v, l2v = 0, l2.Val
			l2 = l2.Next
		}

		curr.Val = l1v + l2v + rip
		if curr.Val >= 10 {
			curr.Val = curr.Val % 10
			rip = 1
		} else {
			rip = 0
		}
		var next ListNode
		curr.Next = &next
		prev = curr
		curr = curr.Next
	}
	if rip != 0 {
		curr.Val += rip
	} else {
		prev.Next = nil
	}
	return &first
}
