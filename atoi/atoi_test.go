package atoi_test

import (
	"leetcode/atoi"
	"testing"
)

func TestAtoi(t *testing.T) {

	testCases := []struct {
		name           string
		str            string
		expectedResult int
	}{
		{
			name:           "simple numeric value",
			str:            "42",
			expectedResult: 42,
		},
		{
			name:           "negative numeric with whitespaces",
			str:            "     -42",
			expectedResult: -42,
		},
		{
			name:           "numeric with words in the end",
			str:            "4193 with words",
			expectedResult: 4193,
		},
		{
			name:           "numerice with some words befoe",
			str:            "words and 987",
			expectedResult: 0,
		},
		{
			name:           "value less than int32 minimum",
			str:            "-91283472332",
			expectedResult: -2147483648,
		},
		{
			name:           "whitespacere after first num",
			str:            "   +0 123",
			expectedResult: 0,
		},
		{
			name:           "numeric with zeros in the beginning",
			str:            "  0000000000012345678",
			expectedResult: 12345678,
		},
		{
			name:           "two signs in a row",
			str:            "+-2",
			expectedResult: 0,
		},
		{
			name:           "two signs in a row part 2",
			str:            "-+1",
			expectedResult: 0,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := atoi.MyAtoi(testCase.str)
			if result != testCase.expectedResult {
				t.Errorf(
					"Failed testcase %s, with input %s:\nexpected %d, got %d",
					testCase.name,
					testCase.str,
					testCase.expectedResult,
					result,
				)
			}
		})
	}
}
