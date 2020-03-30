package mergeklists

import "math"

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func allListNodesAreNil(lists []*ListNode) bool {

	var allListNodesAreNil bool = true
	for _, listNode := range lists {
		if listNode != nil {
			allListNodesAreNil = false
			break
		}
	}

	return allListNodesAreNil
}

func minIndex(vals []int) int {

	var currMin int = math.MaxInt64
	var currMinIndex int

	for ind, val := range vals {
		if val < currMin {
			currMin = val
			currMinIndex = ind
		}
	}

	return currMinIndex
}

func someNodesAreNotNil(lists []*ListNode) bool {

	var someNodesAreNotNil bool
	for _, listNode := range lists {

		if listNode != nil {
			someNodesAreNotNil = true
			break
		}

	}

	return someNodesAreNotNil
}

func mergeKLists(lists []*ListNode) *ListNode {

	if len(lists) == 0 || allListNodesAreNil(lists) {
		return nil
	}

	var result *ListNode
	currVals := make([]int, len(lists))

	for ind := range lists {
		currVals[ind] = math.MaxInt64
	}

	for ind, ln := range lists {
		if ln != nil {
			currVals[ind] = ln.Val
		}
	}

	minValIndex := minIndex(currVals)
	result = &ListNode{
		Val:  currVals[minValIndex],
		Next: nil,
	}

	if lists[minValIndex] != nil {
		lists[minValIndex] = lists[minValIndex].Next
	}

	if someNodesAreNotNil(lists) {
		result.Next = mergeKLists(lists)
	}

	return result
}
