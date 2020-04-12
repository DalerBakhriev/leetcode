package countnumberswithuniquedigits

import (
	"math"
	"strconv"
	"strings"
)

func backTrack(
	solutions map[string]int,
	candidates []int,
	solution string,
	startIndex int,
	targetLength int,
	freqMap map[int]int,
	high int,
) {

	if len(solution) > targetLength {
		return
	}

	if len(solution) == targetLength {

		if strings.HasPrefix(solution, "0") && len(solution) > 1 {
			return
		}

		if solutionAsNum, _ := strconv.Atoi(solution); solutionAsNum >= high {
			return
		}

		solutions[solution] = 1
		return
	}

	for ind := startIndex; ind < len(candidates); ind++ {
		if freqMap[candidates[ind]] > 0 && !strings.Contains(solution, strconv.Itoa(candidates[ind])) {
			freqMap[candidates[ind]]--
			backTrack(
				solutions,
				candidates,
				strings.Join([]string{solution, strconv.Itoa(candidates[ind])}, ""),
				startIndex,
				targetLength,
				freqMap,
				high,
			)
			freqMap[candidates[ind]]++
		}

	}
}

// CountNumbersWithUniqueDigits ...
func CountNumbersWithUniqueDigits(n int) int {

	if n == 0 {
		return 1
	}

	high := int(math.Pow10(n))
	highLength := len(strconv.Itoa(high))

	candidates := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	solutions := make(map[string]int)
	seqLength := 1

	for seqLength >= 1 && seqLength <= highLength {
		freqMap := map[int]int{
			0: 1,
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
		backTrack(solutions, candidates, "", 0, seqLength, freqMap, high)
		seqLength++
	}

	return len(solutions)
}
