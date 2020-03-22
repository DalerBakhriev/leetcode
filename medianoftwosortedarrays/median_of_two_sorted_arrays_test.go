package medianoftwosortedarrays_test

import (
	mtsa "leetcode/medianoftwosortedarrays"
	"reflect"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {

	testCases := []struct {
		name           string
		arrayOne       []int
		arrayTwo       []int
		expectedResult float64
	}{
		{
			name:           "simple test example",
			arrayOne:       []int{1, 3},
			arrayTwo:       []int{2},
			expectedResult: 2.0,
		},
		{
			name:           "not integer median",
			arrayOne:       []int{1, 2},
			arrayTwo:       []int{3, 4},
			expectedResult: 2.5,
		},
		{
			name:           "two simple arrays",
			arrayOne:       []int{1},
			arrayTwo:       []int{1},
			expectedResult: 1.0,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := mtsa.FindMedianSortedArrays(testCase.arrayOne, testCase.arrayTwo)
			if result != testCase.expectedResult {
				t.Errorf(
					"Failed test %s. Expected %f, got %f",
					testCase.name,
					testCase.expectedResult,
					result,
				)
			}
		})
	}

}

func TestMergeSortedArrays(t *testing.T) {

	testCases := []struct {
		name           string
		arrayOne       []int
		arrayTwo       []int
		expectedResult []int
	}{
		{
			name:           "simple test example",
			arrayOne:       []int{1, 3},
			arrayTwo:       []int{2},
			expectedResult: []int{1, 2, 3},
		},
		{
			name:           "not integer median",
			arrayOne:       []int{1, 2},
			arrayTwo:       []int{3, 4},
			expectedResult: []int{1, 2, 3, 4},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := mtsa.MergeSortedArrays(testCase.arrayOne, testCase.arrayTwo)
			if !reflect.DeepEqual(result, testCase.expectedResult) {
				t.Errorf(
					"Failed test %s. Expected %v, got %v",
					testCase.name,
					testCase.expectedResult,
					result,
				)
			}
		})
	}

}
