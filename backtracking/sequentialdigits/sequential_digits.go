package sequentialdigits

import (
	"sort"
	"strconv"
	"strings"
)

func isSequential(nums []int) bool {
	isSequential := true
	for ind := 0; ind < len(nums)-1; ind++ {
		numInt := nums[ind]
		nextNumInt := nums[ind+1]
		if nextNumInt-numInt != 1 {
			isSequential = false
			break
		}
	}
	return isSequential
}

func backTrack(
	solutions *[][]int,
	candidates []int,
	solution []int,
	startIndex int,
	targetLength int,
	freqMap map[int]int,
) {

	if len(solution) > targetLength {
		return
	}

	if len(solution) == targetLength && isSequential(solution) {
		solutionCopy := make([]int, len(solution))
		copy(solutionCopy, solution)
		*solutions = append(*solutions, solutionCopy)
		return
	}

	for ind := startIndex; ind < len(candidates); ind++ {
		if freqMap[candidates[ind]] > 0 {
			freqMap[candidates[ind]]--
			solution = append(solution, candidates[ind])
			backTrack(solutions, candidates, solution, startIndex, targetLength, freqMap)
			freqMap[candidates[ind]]++
			solution = solution[:len(solution)-1]
		}

	}
}

func convertToStringArray(intArray []int) []string {
	stringArr := make([]string, 0)
	for _, num := range intArray {
		stringArr = append(stringArr, strconv.Itoa(num))
	}

	return stringArr
}

// SequentialDigits ...
func SequentialDigits(low int, high int) []int {

	lowLength := len(strconv.Itoa(low))
	highLength := len(strconv.Itoa(high))

	candidates := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	solutions := [][]int{}
	seqLength := lowLength

	for seqLength >= lowLength && seqLength <= highLength {
		freqMap := map[int]int{
			1: 1,
			2: 1,
			3: 1,
			4: 1,
			5: 1,
			6: 1,
			7: 1,
			8: 1,
			9: 1,
		}
		backTrack(&solutions, candidates, []int{}, 0, seqLength, freqMap)
		seqLength++
	}

	finalSolutionsUnique := make(map[int]int, 0)
	for _, solution := range solutions {
		solutionAsStringArr := convertToStringArray(solution)
		solutionAsNum, _ := strconv.Atoi(strings.Join(solutionAsStringArr, ""))
		finalSolutionsUnique[solutionAsNum] = 1
	}

	finalSolutions := make([]int, 0)
	for solution := range finalSolutionsUnique {
		if solution >= low && solution <= high {
			finalSolutions = append(finalSolutions, solution)
		}
	}
	sort.Ints(finalSolutions)

	return finalSolutions
}
