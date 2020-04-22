package permutations

func backTrack(solutions *[][]int, candidates []int, solution []int) {

	if len(candidates) == 0 {
		solutionCopy := make([]int, len(solution))
		copy(solutionCopy, solution)
		*solutions = append(*solutions, solutionCopy)
		return
	}

	for ind, candidate := range candidates {
		solution = append(solution, candidate)
		candidatesCopy := make([]int, len(candidates))
		copy(candidatesCopy, candidates)

		candidatesCopy = append(candidatesCopy[:ind], candidatesCopy[ind+1:]...)
		backTrack(solutions, candidatesCopy, solution)

		solution = solution[:len(solution)-1]
	}

}

// Permute ...
func Permute(nums []int) [][]int {

	var solutions [][]int

	if len(nums) != 0 {
		backTrack(&solutions, nums, []int{})
	}

	return solutions
}
