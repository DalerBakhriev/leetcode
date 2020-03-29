package threesum

import (
	"reflect"
	"sort"
	"strconv"
	"strings"
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

func getIndexWithExclusion(elementsIndexes []int, indexToExclude int) int {

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
func twoSum(nums []int, target int) [][]int {

	arrayElementsMap := makeHashMap(nums)
	var delta int
	var result [][]int

	for index, element := range nums {
		var resultForThisElement []int
		delta = target - element
		availableIndexes, ok := arrayElementsMap[delta]

		if ok && (delta != element || len(availableIndexes) > 1) {
			secondElementIndex := getIndexWithExclusion(availableIndexes, index)
			resultForThisElement = append(resultForThisElement, element)
			resultForThisElement = append(resultForThisElement, nums[secondElementIndex])
			result = append(result, resultForThisElement)
		}
	}

	return result
}

func removeElement(nums []int, indexOfElementToRemove int) []int {

	nums[indexOfElementToRemove] = nums[len(nums)-1]

	return nums[:len(nums)-1]
}

func resultAlreadyContainsSolution(solution []int, result [][]int) bool {

	var answer bool
	for _, oneOfResults := range result {
		if reflect.DeepEqual(oneOfResults, solution) {
			answer = true
			break
		}
	}

	return answer
}

func makeStringForHashMap(answer []int) string {
	answerArr := make([]string, 3)
	for ind, answerNum := range answer {
		answerArr[ind] = strconv.Itoa(answerNum)
	}

	return strings.Join(answerArr, "")
}

// ThreeSum ...
func ThreeSum(nums []int) [][]int {

	usedNums := make(map[int]int)
	usedAnswers := make(map[string]int)
	var result [][]int

	for i := 0; i < len(nums); i++ {

		if _, alreadyUsedThisNum := usedNums[nums[i]]; alreadyUsedThisNum {
			continue
		}

		usedNums[nums[i]] = 1
		arr := make([]int, len(nums))
		copy(arr, nums)
		arrWithoutElement := removeElement(arr, i)
		targetElement := nums[i]

		twoMoreValuesCombinations := twoSum(arrWithoutElement, -1*targetElement)
		if len(twoMoreValuesCombinations) != 0 {

			for _, additionalPair := range twoMoreValuesCombinations {
				additionalPair = append(additionalPair, targetElement)
				sort.Ints(additionalPair)
				answerKey := makeStringForHashMap(additionalPair)

				if _, ok := usedAnswers[answerKey]; !ok {
					result = append(result, additionalPair)
					usedAnswers[answerKey] = 1
				}
			}
		}
	}

	return result
}
