package twosum

import (
	"fmt"
)

func makeHashMap(nums []int) map[int][]int {

	arrayElements := make(map[int][]int, len(nums))
	for index, element := range nums {
		_, ok := arrayElements[element]
		if ok {
			arrayElements[element] = append(arrayElements[element], index)
		} else {
			arrayElements[element] = []int{index}
		}
	}

	return arrayElements
}

func getIndexWithExclusion(element int, elementsIndexes []int, indexToExclude int) int {

	var resultIndex int

	for _, ind := range elementsIndexes {
		if ind != indexToExclude {
			resultIndex = ind
			break
		}
	}

	return resultIndex
}

// FasterWithMap gets pair of ints which adds up to taget value
func FasterWithMap(nums []int, target int) []int {

	arrayElementsMap := makeHashMap(nums)
	fmt.Println(arrayElementsMap)
	var delta int
	result := make([]int, 2)

	for index, element := range nums {
		delta = target - element
		availableIndexes, ok := arrayElementsMap[delta]
		if ok && (delta != element || len(availableIndexes) > 1) {
			secondElementIndex := getIndexWithExclusion(delta, availableIndexes, index)
			result[0] = index
			result[1] = secondElementIndex
			break
		}
	}

	return result
}
