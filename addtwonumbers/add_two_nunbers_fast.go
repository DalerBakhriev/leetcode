package addtwonumbers

func addTwoNums(l1 *ListNode, l2 *ListNode, overflow int) *ListNode {

	var firstNum int
	var secondNum int

	if l1 != nil {
		firstNum = l1.Val
	}

	if l2 != nil {
		secondNum = l2.Val
	}

	sumRes := firstNum + secondNum + overflow
	divRes := sumRes / 10
	divMod := sumRes % 10

	result := &ListNode{
		Val:  divMod,
		Next: nil,
	}

	var newFirstNum *ListNode
	var newSecondNum *ListNode

	if l1 != nil {
		newFirstNum = l1.Next
	}

	if l2 != nil {
		newSecondNum = l2.Next
	}

	if newFirstNum != nil || newSecondNum != nil || divRes > 0 {
		result.Next = addTwoNums(newFirstNum, newSecondNum, divRes)
	}

	return result
}

func addTwoNumbersFast(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNums(l1, l2, 0)
}
