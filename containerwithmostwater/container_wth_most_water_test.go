package containerwithmostwater_test

import (
	cwmw "leetcode/containerwithmostwater"
	"testing"
)

func TestMaxAres(t *testing.T) {

	testCases := []struct {
		name           string
		heights        []int
		expectedResult int
	}{
		{
			name:           "simple heights",
			heights:        []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expectedResult: 49,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := cwmw.MaxArea(testCase.heights)
			if result != testCase.expectedResult {
				t.Errorf(
					"Failed testcase %s with input %v.\nExpected %d, got %d",
					testCase.name,
					testCase.heights,
					testCase.expectedResult,
					result,
				)
			}
		})
	}
}
