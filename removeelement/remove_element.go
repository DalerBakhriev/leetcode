package removeelement

func firstIndexOf(nums []int, val int) int {

	indexOfVal := -1
	for index, element := range nums {
		if element == val {
			indexOfVal = index
			break
		}
	}

	return indexOfVal
}

func removeElementByIndex(nums []int, index int) []int {

	nums[index] = nums[len(nums)-1]

	return nums[:len(nums)-1]
}

func removeElement(nums []int, val int) int {

	indexToRemove := firstIndexOf(nums, val)
	for indexToRemove != -1 {
		nums = removeElementByIndex(nums, indexToRemove)
		indexToRemove = firstIndexOf(nums, val)
	}

	return len(nums)
}
