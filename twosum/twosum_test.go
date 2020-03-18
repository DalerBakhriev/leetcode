package twosum_test

import (
	"leetcode/twosum"
	"reflect"
	"testing"
)

type TwoSumTestCase struct {
	name           string
	nums           []int
	target         int
	expectedResult []int
}

var testCases []TwoSumTestCase = []TwoSumTestCase{
	TwoSumTestCase{
		name:           "two similar values",
		nums:           []int{3, 3},
		target:         6,
		expectedResult: []int{0, 1},
	},
	TwoSumTestCase{
		name:           "simple case",
		nums:           []int{2, 7, 11, 15},
		target:         9,
		expectedResult: []int{0, 1},
	},
}

func TestBruteForce(t *testing.T) {

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := twosum.BruteForce(testCase.nums, testCase.target)
			if !reflect.DeepEqual(result, testCase.expectedResult) {
				t.Errorf(
					"Failed test '%s'. expected %v, got %v",
					testCase.name,
					testCase.expectedResult,
					result,
				)
			}
		})
	}
}

func TestFasterWithMap(t *testing.T) {

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := twosum.FasterWithMap(testCase.nums, testCase.target)
			if !reflect.DeepEqual(result, testCase.expectedResult) {
				t.Errorf(
					"Failed test '%s'. expected %v, got %v",
					testCase.name,
					testCase.expectedResult,
					result,
				)
			}
		})
	}
}
