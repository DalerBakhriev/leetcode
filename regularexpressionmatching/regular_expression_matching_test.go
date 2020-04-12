package regularexpressionmatching_test

import (
	rem "leetcode/regularexpressionmatching"
	"testing"
)

type TestCase struct {
	name           string
	s              string
	p              string
	expectedResult bool
}

var testCases []TestCase = []TestCase{
	{
		name:           "entire string is not matched",
		s:              "aa",
		p:              "a",
		expectedResult: false,
	},
	{
		name:           "entire string matches with *",
		s:              "aa",
		p:              "a*",
		expectedResult: true,
	},
	{
		name:           "zero or more of any character",
		s:              "ab",
		p:              ".*",
		expectedResult: true,
	},
	{
		name:           "match with *",
		s:              "aab",
		p:              "c*a*b",
		expectedResult: true,
	},
	{
		name:           "not matched mississippi",
		s:              "mississippi",
		p:              "mis*is*p*.",
		expectedResult: false,
	},
	{
		name:           "matched mississipi",
		s:              "mississippi",
		p:              "mis*is*ip*.",
		expectedResult: true,
	},
	{
		name:           "string shorter than pattern",
		s:              "ab",
		p:              ".*c",
		expectedResult: false,
	},
	{
		name:           "string shorter than pattern one more",
		s:              "aaa",
		p:              "ab*ac*a",
		expectedResult: true,
	},
	{
		name:           "short",
		s:              "abcd",
		p:              "d*",
		expectedResult: false,
	},
	{
		name:           "'a's",
		s:              "aaa",
		p:              "ab*a",
		expectedResult: false,
	},
	{
		name:           "one more",
		s:              "aaca",
		p:              "ab*a*c*a",
		expectedResult: true,
	},
	{
		name:           "simple with one symbol in s",
		s:              "a",
		p:              "ab*a",
		expectedResult: false,
	},
	{
		name:           "blyaaaaaa",
		s:              "aaa",
		p:              "ab*a*c*a",
		expectedResult: true,
	},
	{
		name:           "simple one more",
		s:              "a",
		p:              "ab*",
		expectedResult: true,
	},
	{
		name:           "fucking a`s",
		s:              "aaa",
		p:              "a*a",
		expectedResult: true,
	},
	{
		name:           "b`s",
		s:              "bbbba",
		p:              ".*a*a",
		expectedResult: true,
	},
}

func TestIsMatch(t *testing.T) {

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := rem.IsMatch(testCase.s, testCase.p)
			if result != testCase.expectedResult {
				t.Errorf(
					"Failed testscase %s, with inputs\n s: %s, p: %s.\n Expected %v, got %v",
					testCase.name,
					testCase.s,
					testCase.p,
					testCase.expectedResult,
					result,
				)
			}
		})
	}
}

func TestIsMatchMemo(t *testing.T) {

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := rem.IsMatch(testCase.s, testCase.p)
			if result != testCase.expectedResult {
				t.Errorf(
					"Failed testscase %s, with inputs\n s: %s, p: %s.\n Expected %v, got %v",
					testCase.name,
					testCase.s,
					testCase.p,
					testCase.expectedResult,
					result,
				)
			}
		})
	}
}
