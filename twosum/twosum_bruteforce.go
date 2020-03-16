package twosum

func remove(slice []int, indexOfElementToRemove int) []int {

	sliceAfterRemoving := make([]int, len(slice))
	copy(sliceAfterRemoving, slice)
	sliceAfterRemoving = append(
		sliceAfterRemoving[:indexOfElementToRemove],
		sliceAfterRemoving[indexOfElementToRemove+1:]...,
	)
	return sliceAfterRemoving
}

func indexOf(element int, data []int, excludeIndex int) int {

	for index, value := range data {
		if element == value && index != excludeIndex {
			return index
		}
	}
	return -1 //not found.
}

// BruteForce gets pair of ints which adds up to taget value
func BruteForce(nums []int, target int) []int {

	result := []int{0, 0}
	var solutionFound bool

	for indexFirst, numFirst := range nums {

		numsWithoutFirstElement := remove(nums, indexFirst)

		for _, numSecond := range numsWithoutFirstElement {
			if numSecond+numFirst == target {
				result[0] = indexFirst
				result[1] = indexOf(numSecond, nums, indexFirst)
				solutionFound = true
				break
			}
		}

		if solutionFound {
			break
		}
	}

	return result
}
