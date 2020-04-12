package combinationsum

import (
	"reflect"
	"sort"
)

var solutions [][]int

func clearSolutions() {
	solutions = make([][]int, 0)
}

// CombinationSum ...
func CombinationSum(candidates []int, target int) [][]int {

	defer clearSolutions()
	var solution []int
	backTrack(candidates, target, solution, 0)

	return solutions
}

func contains(solutions [][]int, solution []int) bool {
	var contains bool

	for _, solutionFromSolutions := range solutions {
		if reflect.DeepEqual(solution, solutionFromSolutions) {
			contains = true
			break
		}
	}

	return contains
}

func backTrack(candidates []int, remains int, solution []int, startIndex int) {

	if remains < 0 {
		return
	} else if remains == 0 {
		solutionToAppend := make([]int, len(solution))
		copy(solutionToAppend, solution)
		sort.Ints(solutionToAppend)
		if !contains(solutions, solutionToAppend) {
			solutions = append(solutions, solutionToAppend)
			return
		}

	} else {
		for ind := range candidates {
			solution = append(solution, candidates[ind])
			backTrack(candidates, remains-candidates[ind], solution, ind)
			solution = solution[:len(solution)-1]
		}
	}
}
