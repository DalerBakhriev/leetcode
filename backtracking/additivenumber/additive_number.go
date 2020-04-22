package additivenumber

import (
	"strconv"
)

func backTrack(sourceNumber string, solution []string, startIndex int) bool {

	if startIndex == len(sourceNumber) && len(solution) > 2 {
		return true
	}

	result := false
	for ind := startIndex; ind < len(sourceNumber); ind++ {
		currNum := sourceNumber[startIndex : ind+1]

		if len(currNum) > 1 && currNum[0] == '0' {
			continue
		}

		if len(solution) >= 2 {

			firstElement, _ := strconv.Atoi(solution[len(solution)-2])
			secondElement, _ := strconv.Atoi(solution[len(solution)-1])
			currNumAsInt, _ := strconv.Atoi(currNum)

			if firstElement+secondElement != currNumAsInt {
				continue
			}
		}

		solution = append(solution, currNum)

		result = result || backTrack(sourceNumber, solution, ind+1)
		solution = solution[:len(solution)-1]
	}
	return result
}

// IsAdditiveNumber ...
func IsAdditiveNumber(num string) bool {
	return backTrack(num, []string{}, 0)
}
