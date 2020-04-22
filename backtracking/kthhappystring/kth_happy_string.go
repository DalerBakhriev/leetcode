package kthhapppystring

import (
	"sort"
)

func backTrack(happyStrings map[string]int, candidates []string, happyString string, happyStringLen int, happyStringNum int) {

	if len(happyStrings) > happyStringNum {
		return
	}

	if len(happyString) == happyStringLen {
		if _, happyStringExists := happyStrings[happyString]; !happyStringExists {
			happyStrings[happyString] = 1
		}
		return
	}

	for _, candidate := range candidates {

		if len(happyString) > 0 && candidate == string(happyString[len(happyString)-1]) {
			continue
		}
		backTrack(happyStrings, candidates, happyString+string(candidate), happyStringLen, happyStringNum)
	}

}

// GetHappyString ...
func GetHappyString(n int, k int) string {

	allPossibleLetters := []string{"a", "b", "c"}
	numRepeatTimes := (n / len(allPossibleLetters)) + 1
	candidatesLetters := make([]string, 0)

	for i := 0; i < numRepeatTimes; i++ {
		candidatesLetters = append(candidatesLetters, allPossibleLetters...)
	}

	happyStrings := make(map[string]int)
	backTrack(happyStrings, candidatesLetters, "", n, k)
	if k > len(happyStrings) {
		return ""
	}

	happyStringsArr := make([]string, 0)
	for happyString := range happyStrings {
		happyStringsArr = append(happyStringsArr, happyString)
	}
	sort.Strings(happyStringsArr)

	return happyStringsArr[k-1]
}
