package longestsubstringwithoutrepeating_test

import (
	lswr "leetcode/longestsubstringwithoutrepeating"
	"testing"
)

type TestCase struct {
	name           string
	testString     string
	expectedResult int
}

var testCases = []TestCase{
	TestCase{
		name:           "normal length string",
		testString:     "abcabcbb",
		expectedResult: 3,
	},
	TestCase{
		name:           "one character string",
		testString:     "bbbbb",
		expectedResult: 1,
	},
	TestCase{
		name:           "substring in the middle",
		testString:     "pwwkew",
		expectedResult: 3,
	},
	TestCase{
		name:           "empty string",
		testString:     " ",
		expectedResult: 1,
	},
	TestCase{
		name:           "sliding windows substring",
		testString:     "dvdf",
		expectedResult: 3,
	},
	TestCase{
		name:           "bigger sliding windows",
		testString:     "dtvdef",
		expectedResult: 5,
	},
	TestCase{
		name:           "simple testcase",
		testString:     "cdd",
		expectedResult: 2,
	},
	TestCase{
		name:           "abba",
		testString:     "abba",
		expectedResult: 2,
	},
}

func TestLengthOfLongestSubstring(t *testing.T) {

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := lswr.LengthOfLongestSubstring(testCase.testString)
			if result != testCase.expectedResult {
				t.Errorf(
					"Failed testcase %s, expected %d, found %d",
					testCase.name,
					testCase.expectedResult,
					result,
				)
			}
		})
	}
}

func TestLengthOfLongestSubstringFast(t *testing.T) {

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := lswr.LengthOfLongestSubstringFast(testCase.testString)
			if result != testCase.expectedResult {
				t.Errorf(
					"Failed testcase %s, expected %d, found %d",
					testCase.name,
					testCase.expectedResult,
					result,
				)
			}
		})
	}
}
