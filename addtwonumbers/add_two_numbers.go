package addtwonumbers

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

// ListNode represents nonzero integer number
type ListNode struct {
	Val  int
	Next *ListNode
}

func convertLinkedListToInt(node *ListNode) *big.Int {

	currNode := node

	numAsStringSlice := []string{strconv.Itoa(currNode.Val)}
	for currNode.Next != nil {
		numAsStringSlice = append(numAsStringSlice, strconv.Itoa(currNode.Next.Val))
		currNode = currNode.Next
	}

	numAsString := strings.Join(numAsStringSlice, "")
	revertedNumAsString := reverse(numAsString)

	resultNum, _ := strconv.Atoi(revertedNumAsString)

	return big.NewInt(int64(resultNum))
}

func reverse(stringToReverse string) string {

	var revertedstringSlice []string
	for ind := len(stringToReverse) - 1; ind >= 0; ind-- {
		revertedstringSlice = append(revertedstringSlice, string(stringToReverse[ind]))
	}

	return strings.Join(revertedstringSlice, "")
}

func convertIntSliceToLinkedList(intSlice []int, elementsIndex int) *ListNode {

	if elementsIndex == len(intSlice) {
		return nil
	}

	return &ListNode{
		Val:  intSlice[elementsIndex],
		Next: convertIntSliceToLinkedList(intSlice, elementsIndex+1),
	}
}

func convertStringToLinkedList(numAsString string) *ListNode {

	var numAsIntSlice []int

	for ind := len(numAsString) - 1; ind >= 0; ind-- {
		element, _ := strconv.Atoi(string(numAsString[ind]))
		numAsIntSlice = append(numAsIntSlice, element)
	}

	return convertIntSliceToLinkedList(numAsIntSlice, 0)
}

// AddTwoNumbers adds up two nonzero integers represented as listnode
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	listOneAsNum := convertLinkedListToInt(l1)
	listTwoAsNum := convertLinkedListToInt(l2)

	resultAsBigInt := new(big.Int)
	resultAsBigInt.Add(listOneAsNum, listTwoAsNum)
	fmt.Printf("First list as number %d\n", listOneAsNum)
	fmt.Printf("Second list as number %d\n", listTwoAsNum)
	fmt.Printf("Result as number %v\n", resultAsBigInt)

	return convertStringToLinkedList(resultAsBigInt.String())
}
