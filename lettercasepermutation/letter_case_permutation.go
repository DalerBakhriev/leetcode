package lettercasepermutation

import (
	"strconv"
	"strings"
)

func backTrack(
	solutions map[string]int,
	candidates string,
	solution string,
	currIndex int,
) {

	if len(solution) > len(candidates) {
		return
	}

	if len(solution) == len(candidates) && strings.ToLower(solution) == strings.ToLower(candidates) {
		if _, ok := solutions[solution]; !ok {
			solutions[solution] = 1
		}
		return
	}

	for _, shouldCapitalize := range []bool{true, false} {
		candidateToAppend := string(candidates[currIndex])
		if _, err := strconv.Atoi(candidateToAppend); err != nil {
			if shouldCapitalize {
				candidateToAppend = strings.ToUpper(candidateToAppend)
			} else {
				candidateToAppend = strings.ToLower(candidateToAppend)
			}
		}
		backTrack(
			solutions,
			candidates,
			strings.Join([]string{solution, candidateToAppend}, ""),
			currIndex+1,
		)
	}

}

// LetterCasePermutation ...
func LetterCasePermutation(S string) []string {

	solutionsUnique := make(map[string]int, 0)
	backTrack(solutionsUnique, S, "", 0)

	solutions := make([]string, 0)

	for solution := range solutionsUnique {
		solutions = append(solutions, solution)
	}

	return solutions
}
