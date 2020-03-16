package reverseinteger

import (
	"math"
	"strconv"
	"strings"
)

// Reverse inverses integer number
func Reverse(numberToReverse int) int {

	numberAsString := strconv.Itoa(numberToReverse)
	var resultAsSlice []string
	var endIndex int

	if strings.HasPrefix(numberAsString, "-") {
		endIndex = 1
		resultAsSlice = append(resultAsSlice, "-")
	}

	for ind := len(numberAsString) - 1; ind >= endIndex; ind-- {
		resultAsSlice = append(resultAsSlice, string(numberAsString[ind]))
	}

	result, _ := strconv.Atoi(strings.Join(resultAsSlice, ""))
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}

	return result
}
