package mergetwosortedlists

import "math"

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// MergeTwoSortedLists ...
func MergeTwoSortedLists(l1 *ListNode, l2 *ListNode) *ListNode {

	var currFirstVal int = math.MaxInt64
	var currSecondVal int = math.MaxInt64
	var nextSecondNode *ListNode
	var nextFirstNode *ListNode

	var result *ListNode

	if l1 != nil {
		currFirstVal = l1.Val
	}

	if l2 != nil {
		currSecondVal = l2.Val
	}

	if l1 == nil && l2 == nil {
		return nil
	}

	if currFirstVal < currSecondVal {
		result = &ListNode{Val: currFirstVal, Next: nil}
		if l1 != nil {
			nextFirstNode = l1.Next
		}
		nextSecondNode = l2
	} else {
		result = &ListNode{Val: currSecondVal, Next: nil}
		if l2 != nil {
			nextFirstNode = l2.Next
		}
		nextSecondNode = l1
	}

	if nextFirstNode != nil || nextSecondNode != nil {
		result.Next = MergeTwoSortedLists(nextFirstNode, nextSecondNode)
	}

	return result
}
