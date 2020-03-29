package longestpalindrome_test

import (
	lp "leetcode/longestpalindrome"
	"testing"
)

type TestCase struct {
	name           string
	s              string
	expectedResult []string
}

var testCases []TestCase = []TestCase{
	TestCase{
		name:           "simple palindrome substring",
		s:              "babad",
		expectedResult: []string{"bab", "aba"},
	},
	TestCase{
		name:           "repeating character",
		s:              "cbbd",
		expectedResult: []string{"bb"},
	},
	TestCase{
		name:           "single letter palindrome",
		s:              "ac",
		expectedResult: []string{"a", "c"},
	},
	TestCase{
		name:           "empty string",
		s:              "",
		expectedResult: []string{""},
	},
	TestCase{
		name:           "repeating character in the end",
		s:              "abb",
		expectedResult: []string{"bb"},
	},
}

func arrayContains(array []string, s string) bool {

	var result bool
	for _, element := range array {
		if element == s {
			result = true
			break
		}
	}

	return result
}

func TestLongestPalindromeBruteForce(t *testing.T) {

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := lp.BruteForce(testCase.s)
			if !arrayContains(testCase.expectedResult, result) {
				t.Errorf(
					"Failed testcase %s, with input %s:\nexpected one of %v, got %s",
					testCase.name,
					testCase.s,
					testCase.expectedResult,
					result,
				)
			}
		})
	}
}

func TestLongestPalindrome(t *testing.T) {

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := lp.LongestPalindrome(testCase.s)
			if !arrayContains(testCase.expectedResult, result) {
				t.Errorf(
					"Failed testcase %s, with input %s:\nexpected one of %v, got %s",
					testCase.name,
					testCase.s,
					testCase.expectedResult,
					result,
				)
			}
		})
	}
}
