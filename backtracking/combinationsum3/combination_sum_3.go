package combinationsum3

func backTrack(
	solutions *[][]int,
	candidates []int,
	remains int,
	solutionLength int,
	solution []int,
	startIndex int,
	freqMap map[int]int,
) {

	if remains < 0 {
		return
	} else if remains == 0 {
		if len(solution) != solutionLength {
			return
		}
		solutionToAppend := make([]int, len(solution))
		copy(solutionToAppend, solution)
		*solutions = append(*solutions, solutionToAppend)
		return
	} else {
		for ind := startIndex; ind < len(candidates); ind++ {
			if freqMap[candidates[ind]] > 0 {
				freqMap[candidates[ind]]--
				solution = append(solution, candidates[ind])
				backTrack(solutions, candidates, remains-candidates[ind], solutionLength, solution, ind, freqMap)
				freqMap[candidates[ind]]++
				solution = solution[:len(solution)-1]
			}

		}
	}
}

// CombinationSum3 ...
func CombinationSum3(k int, n int) [][]int {
	candidates := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
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
	var solutions [][]int
	backTrack(&solutions, candidates, n, k, []int{}, 0, freqMap)

	return solutions
}
