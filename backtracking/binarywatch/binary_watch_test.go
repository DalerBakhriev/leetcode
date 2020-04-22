package binarywatch_test

import (
	"leetcode/binarywatch"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func contains(arrToCheck []string, valToCheck string) bool {

	contains := false
	for _, val := range arrToCheck {
		if val == valToCheck {
			contains = true
			break
		}
	}
	return contains
}

// TestReadBinaryWatch ...
func TestReadBinaryWatch(t *testing.T) {

	for i := 0; i < 6; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			myRes := binarywatch.ReadBinaryWatch(i)
			corrRes := binarywatch.ReadBinaryWatchCorrect(i)
			sort.Strings(myRes)
			sort.Strings(corrRes)

			if !reflect.DeepEqual(myRes, corrRes) {

				timesInMyResWhichAreNotInCorrectRes := make([]string, 0)
				for _, myTime := range myRes {
					if !contains(corrRes, myTime) {
						timesInMyResWhichAreNotInCorrectRes = append(
							timesInMyResWhichAreNotInCorrectRes,
							myTime,
						)
					}
				}

				timesInCorrResWhichAreNotInMyRes := make([]string, 0)
				for _, corrTime := range corrRes {
					if !contains(myRes, corrTime) {
						timesInCorrResWhichAreNotInMyRes = append(
							timesInCorrResWhichAreNotInMyRes,
							corrTime,
						)
					}
				}

				t.Errorf(
					"Failed test with num %d\n"+
						"Times in my calculation which are not in correct result: %v\n"+
						"Times in correct result which are not in my calculation: %v",
					i,
					timesInMyResWhichAreNotInCorrectRes,
					timesInCorrResWhichAreNotInMyRes,
				)
			}
		})
	}
}
