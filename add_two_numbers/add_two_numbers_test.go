package add_two_numbers

import (
	"slices"
	"testing"
)

type TestCase struct {
	l1       ListArray
	l2       ListArray
	expected ListArray
}

type ListArray []int

func (la ListArray) toListNode() *ListNode {
	var first ListNode
	curr := &first
	for i, v := range la {
		curr.Val = v
		if i != len(la)-1 {
			var next ListNode
			curr.Next = &next
			curr = curr.Next
		}
	}
	return &first
}

func (ln ListNode) toListArray() ListArray {
	var curr *ListNode = &ln
	var la ListArray
	for curr != nil {
		la = append(la, curr.Val)
		curr = curr.Next
	}
	return la
}

func TestAddTwoNumbers(t *testing.T) {
	testCases := []TestCase{
		{ListArray{4}, ListArray{4}, ListArray{8}},
		{ListArray{2, 4, 3}, ListArray{5, 6, 4}, ListArray{7, 0, 8}},
		{ListArray{0}, ListArray{0}, ListArray{0}},
		{ListArray{9, 9, 9, 9, 9, 9, 9}, ListArray{9, 9, 9, 9}, ListArray{8, 9, 9, 9, 0, 0, 0, 1}},
	}
	for _, tc := range testCases {
		actual := addTwoNumbers(tc.l1.toListNode(), tc.l2.toListNode())
		if !slices.Equal(actual.toListArray(), tc.expected) {
			t.Errorf("addTwoNumbers(%v, %v) = %v; expected %v", tc.l1, tc.l2, actual, tc.expected)
		}
	}
}
