package palindromepartitioning2

import "math"

func isPalindrome(s string) bool {

	isPalindrome := true
	start := 0
	end := len(s) - 1

	for start < end {
		if s[start] != s[end] {
			isPalindrome = false
			break
		}
		end--
		start++
	}

	return isPalindrome
}

func backTrack(minCuts *int, solution []string, targetWord string, startIndex int) {

	if len(solution)-1 > *minCuts {
		return
	}

	if startIndex == len(targetWord) {
		if *minCuts > len(solution)-1 {
			*minCuts = len(solution) - 1
		}
		return
	}

	for ind := startIndex; ind < len(targetWord); ind++ {
		if isPalindrome(targetWord[startIndex : ind+1]) {
			solution = append(solution, targetWord[startIndex:ind+1])
			backTrack(minCuts, solution, targetWord, ind+1)
			solution = solution[:len(solution)-1]
		}
	}

}

// Partition ...
func MinCut(s string) int {

	minCutsNumber := math.MaxInt64
	backTrack(&minCutsNumber, []string{}, s, 0)

	return minCutsNumber
}
