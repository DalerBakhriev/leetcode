package palindrome

import (
	"strconv"
	"strings"
)

func invertString(stringToInvert string) string {

	var stringsArr []string
	for ind := len(stringToInvert) - 1; ind >= 0; ind-- {
		stringsArr = append(stringsArr, string(stringToInvert[ind]))
	}

	return strings.Join(stringsArr, "")
}

// IsPalindrome checks if number is palindrome or not
func IsPalindrome(number int) bool {

	if number < 0 {
		return false
	}

	numberAsString := strconv.Itoa(number)
	invertedNumberAsString := invertString(numberAsString)

	invertedNumber, _ := strconv.Atoi(invertedNumberAsString)

	return invertedNumber == number
}
