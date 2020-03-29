package longestpalindrome

import (
	"strings"
)

func isPalindrome(s string) bool {

	var reversedArr []string
	for ind := len(s) - 1; ind >= 0; ind-- {
		reversedArr = append(reversedArr, string(s[ind]))
	}

	reversedS := strings.Join(reversedArr, "")

	return s == reversedS
}

// BruteForce ...
func BruteForce(s string) string {

	if len(s) <= 1 {
		return s
	}

	longestPalindrome := string(s[0])
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			palindromeCandidate := s[j : i+1]
			if (isPalindrome(palindromeCandidate)) && (len(palindromeCandidate) > len(longestPalindrome)) {
				longestPalindrome = palindromeCandidate
			}
		}
	}

	return longestPalindrome

}

func expandAroundCenter(s string, left, right int) int {

	for (left >= 0) && (right < len(s)) && s[left] == s[right] {
		left--
		right++
	}

	return right - left - 1
}

func max(a, b int) int {

	if a > b {
		return a
	}
	return b
}

// LongestPalindrome ...
func LongestPalindrome(s string) string {

	if len(s) <= 1 {
		return s
	}

	var start, end int
	for i := 0; i < len(s); i++ {

		lenOne := expandAroundCenter(s, i, i)
		lenTwo := expandAroundCenter(s, i, i+1)
		length := max(lenOne, lenTwo)

		if length > end-start {
			start = i - (length-1)/2
			end = i + length/2
		}
	}

	return s[start : end+1]
}
