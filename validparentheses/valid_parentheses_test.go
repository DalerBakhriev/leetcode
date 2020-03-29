package validparentheses_test

import (
	vp "leetcode/validparentheses"
	"testing"
)

func TestIsValid(t *testing.T) {

	testCases := []struct {
		name           string
		s              string
		expectedResult bool
	}{
		{
			name:           "simple round brackets",
			s:              "()",
			expectedResult: true,
		},
		{
			name:           "three brackets pairs",
			s:              "()[]{}",
			expectedResult: true,
		},
		{
			name:           "different open and close brackets types",
			s:              "(]",
			expectedResult: false,
		},
		{
			name:           "invalid bracket close order",
			s:              "([)]",
			expectedResult: false,
		},
		{
			name:           "nested brackets correct order",
			s:              "{[]}",
			expectedResult: true,
		},
		{
			name:           "single open bracket",
			s:              "[",
			expectedResult: false,
		},
		{
			name:           "single close bracket",
			s:              "]",
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := vp.IsValid(testCase.s)
			if result != testCase.expectedResult {
				t.Errorf(
					"Failed testcase %s, expected %v, got %v",
					testCase.name,
					testCase.expectedResult,
					result,
				)
			}
		})
	}
}
