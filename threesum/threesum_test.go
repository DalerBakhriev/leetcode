package threesum_test

import (
	"leetcode/threesum"
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {

	testCases := []struct {
		name           string
		nums           []int
		expectedResult [][]int
	}{
		{
			name: "simple testcase",
			nums: []int{-1, 0, 1, 2, -1, -4},
			expectedResult: [][]int{
				{-1, 0, 1},
				{-1, -1, 2},
			},
		},
		{
			name: "many repeated nums",
			nums: []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6},
			expectedResult: [][]int{
				{-4, -2, 6},
				{-4, 0, 4},
				{-4, 1, 3},
				{-4, 2, 2},
				{-2, -2, 4},
				{-2, 0, 2},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := threesum.ThreeSum(testCase.nums)
			if !reflect.DeepEqual(result, testCase.expectedResult) {
				t.Errorf(
					"Failed test %s with input %v.\n Expected %v, got %v",
					testCase.name,
					testCase.nums,
					testCase.expectedResult,
					result,
				)
			}
		})
	}
}
