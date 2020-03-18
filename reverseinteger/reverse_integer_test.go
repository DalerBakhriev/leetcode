package reverseinteger_test

import (
	"leetcode/reverseinteger"
	"testing"
)

func TestReverse(t *testing.T) {

	testCases := []struct {
		name           string
		numberToRevert int
		expectedResult int
	}{
		{
			name:           "negative number",
			numberToRevert: -123,
			expectedResult: -321,
		},
		{
			name:           "positive number",
			numberToRevert: 675,
			expectedResult: 576,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := reverseinteger.Reverse(testCase.numberToRevert)
			if result != testCase.expectedResult {
				t.Errorf(
					"Failed test case: '%s', expected %d, got %d",
					testCase.name,
					testCase.expectedResult,
					result,
				)
			}
		})
	}
}
