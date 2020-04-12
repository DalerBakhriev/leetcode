package permutationsequence

import (
	"strconv"
	"strings"
)

func backTrack(solutions *[]string, candidates []string, combination string, solutionsNum int, combinationLength int) {

	if len(*solutions) == solutionsNum {
		return
	}

	if len(combination) == combinationLength {
		*solutions = append(*solutions, combination)
		return
	}

	for ind, candidate := range candidates {
		combination = strings.Join([]string{combination, candidate}, "")
		candidatesCopy := make([]string, len(candidates))
		copy(candidatesCopy, candidates)
		candidatesCopy = append(candidatesCopy[:ind], candidatesCopy[ind+1:]...)
		backTrack(solutions, candidatesCopy, combination, solutionsNum, combinationLength)
		combination = combination[:len(combination)-1]
	}
}

func getCandidates(n int) []string {

	candidates := make([]string, 0)
	for i := 1; i < n+1; i++ {
		candidates = append(candidates, strconv.Itoa(i))
	}

	return candidates
}

// GetPermutation ...
func GetPermutation(n int, k int) string {

	solutions := make([]string, 0)
	candidates := getCandidates(n)

	backTrack(&solutions, candidates, "", k, n)

	return solutions[k-1]
}
