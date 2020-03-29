package atoi

import (
	"math"
	"strconv"
	"strings"
)

// MyAtoi ...
func MyAtoi(str string) int {

	var resultArr []string
	var result int
	var startIndex int
	var foundFirstNum bool

	for ind, s := range str[startIndex:] {

		char := string(s)
		if char == " " && !foundFirstNum {
			continue
		}

		if char == " " && foundFirstNum {
			break
		}

		if char == "-" && !foundFirstNum {
			resultArr = append(resultArr, string(s))
			foundFirstNum = true
			continue
		}

		if char == "+" && !foundFirstNum {
			foundFirstNum = true
			continue
		}

		if char == "0" && !foundFirstNum {
			foundFirstNum = true
			continue
		}

		_, err := strconv.Atoi(char)
		if err != nil {

			if ind == 0 {
				return result
			}
			break
		}

		foundFirstNum = true
		resultArr = append(resultArr, string(s))

	}

	resultAsString := strings.Join(resultArr, "")
	resultAsInt, _ := strconv.Atoi(resultAsString)

	if resultAsInt > math.MaxInt32 {
		return math.MaxInt32
	}

	if resultAsInt < math.MinInt32 {
		return math.MinInt32
	}

	return resultAsInt
}
