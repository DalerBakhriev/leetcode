package permutations2

func backTrack(solutions *[][]int, candidates []int, solution []int, freqMap map[int]int, uniqueNums []int, elementsNum int) {

	if len(solution) == elementsNum {
		solutionCopy := make([]int, len(solution))
		copy(solutionCopy, solution)
		*solutions = append(*solutions, solutionCopy)
		return
	}

	for _, candidate := range uniqueNums {
		if freqMap[candidate] > 0 {
			freqMap[candidate]--
			backTrack(solutions, candidates, append(solution, candidate), freqMap, uniqueNums, elementsNum)
			freqMap[candidate]++
		}
	}
}

// PermuteUnique ...
func PermuteUnique(nums []int) [][]int {

	var solutions [][]int
	freqMap := make(map[int]int)
	uniqueNums := make([]int, 0)

	for _, num := range nums {
		if _, ok := freqMap[num]; !ok {
			uniqueNums = append(uniqueNums, num)
			freqMap[num] = 1
		} else {
			freqMap[num]++
		}
	}

	backTrack(&solutions, nums, []int{}, freqMap, uniqueNums, len(nums))

	return solutions
}
