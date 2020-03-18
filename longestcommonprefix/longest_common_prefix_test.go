package longestcommonprefix_test

import (
	lcp "leetcode/longestcommonprefix"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {

	testCases := []struct {
		name           string
		strs           []string
		expectedResult string
	}{
		{
			name:           "has common prefix",
			strs:           []string{"flower", "flow", "flight"},
			expectedResult: "fl",
		},
		{
			name:           "has no common prefix",
			strs:           []string{"dog", "racecar", "car"},
			expectedResult: "",
		},
		{
			name:           "empty strings arrays",
			strs:           []string{},
			expectedResult: "",
		},
		{
			name:           "single element strings array",
			strs:           []string{"a"},
			expectedResult: "a",
		},
		{
			name:           "two similar elements in strings array",
			strs:           []string{"c", "c"},
			expectedResult: "c",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := lcp.GetLongestCommonPrefix(testCase.strs)
			if result != testCase.expectedResult {
				t.Errorf(
					"Failed test '%s', expected %s, got %s",
					testCase.name,
					testCase.expectedResult,
					result,
				)
			}
		})
	}
}
