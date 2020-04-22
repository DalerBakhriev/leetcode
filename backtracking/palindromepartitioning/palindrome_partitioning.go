package palindromepartitioning

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

func backTrack(partitions *[][]string, solution []string, targetWord string, startIndex int) {

	if startIndex == len(targetWord) {
		solutionCopy := make([]string, len(solution))
		copy(solutionCopy, solution)
		*partitions = append(*partitions, solutionCopy)
		return
	}

	for ind := startIndex; ind < len(targetWord); ind++ {
		if isPalindrome(targetWord[startIndex : ind+1]) {
			solution = append(solution, targetWord[startIndex:ind+1])
			backTrack(partitions, solution, targetWord, ind+1)
			solution = solution[:len(solution)-1]
		}
	}

}

// Partition ...
func Partition(s string) [][]string {

	partitions := make([][]string, 0)
	backTrack(&partitions, []string{}, s, 0)

	return partitions
}
