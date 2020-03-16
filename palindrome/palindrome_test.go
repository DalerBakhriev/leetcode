package palindrome_test

import (
	"leetcode/palindrome"
	"testing"
)

func TestIsPalindrome(t *testing.T) {

	testCases := []struct {
		name           string
		testNumber     int
		expectedResult bool
	}{
		{
			name:           "negative number",
			testNumber:     -121,
			expectedResult: false,
		},
		{
			name:           "not a palindrome",
			testNumber:     10,
			expectedResult: false,
		},
		{
			name:           "palindrome",
			testNumber:     121,
			expectedResult: true,
		},
	}

	for _, testCase := range testCases {
		result := palindrome.IsPalindrome(testCase.testNumber)
		if result != testCase.expectedResult {
			t.Errorf("Expected %v, got %v", testCase.expectedResult, result)
		}
	}
}
