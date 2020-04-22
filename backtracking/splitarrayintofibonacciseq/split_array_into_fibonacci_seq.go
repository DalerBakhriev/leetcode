package splitarrayintofibonacciseq

import (
	"math"
	"strconv"
)

func backTrack(solutions *[][]string, sourceNumber string, solution []string, startIndex int) {

	if startIndex == len(sourceNumber) && len(solution) > 2 {
		*solutions = append(*solutions, solution)
	}

	for ind := startIndex; ind < len(sourceNumber); ind++ {
		currNum := sourceNumber[startIndex : ind+1]

		if len(currNum) > 1 && currNum[0] == '0' {
			continue
		}
		currNumAsInt, _ := strconv.Atoi(currNum)
		if currNumAsInt > math.MaxInt32 {
			continue
		}

		if len(solution) >= 2 {

			firstElement, _ := strconv.Atoi(solution[len(solution)-2])
			secondElement, _ := strconv.Atoi(solution[len(solution)-1])

			if int32(firstElement)+int32(secondElement) != int32(currNumAsInt) || currNumAsInt > math.MaxInt32 {
				continue
			}
		}

		solution = append(solution, currNum)
		backTrack(solutions, sourceNumber, solution, ind+1)
		solution = solution[:len(solution)-1]
	}
}

// SplitIntoFibonacci ...
func SplitIntoFibonacci(S string) []int {

	solutions := make([][]string, 0)
	backTrack(&solutions, S, []string{}, 0)

	result := make([]int, 0)
	if len(solutions) == 0 {
		return result
	}

	resultWithStrings := solutions[0]
	for _, num := range resultWithStrings {
		numAsInt, _ := strconv.Atoi(string(num))
		result = append(result, numAsInt)
	}

	return result
}
