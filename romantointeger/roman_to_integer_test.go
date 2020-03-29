package romantointeger_test

import (
	rti "leetcode/romantointeger"
	"testing"
)

func TestRomanToInteger(t *testing.T) {

	testCases := []struct {
		name           string
		roman          string
		expectedResult int
	}{
		{
			name:           "three",
			roman:          "III",
			expectedResult: 3,
		},
		{
			name:           "four",
			roman:          "IV",
			expectedResult: 4,
		},
		{
			name:           "nine",
			roman:          "IX",
			expectedResult: 9,
		},
		{
			name:           "fifty eight",
			roman:          "LVIII",
			expectedResult: 58,
		},
		{
			name:           "nineteen ninety four",
			roman:          "MCMXCIV",
			expectedResult: 1994,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := rti.RomanToInt(testCase.roman)
			if result != testCase.expectedResult {
				t.Errorf(
					"Failed test %s. Expected %d, got %d",
					testCase.name,
					testCase.expectedResult,
					result,
				)
			}
		})
	}
}
