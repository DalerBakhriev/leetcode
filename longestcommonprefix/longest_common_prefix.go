package longestcommonprefix

import (
	"math"
	"sort"
)

func getMinWordLength(strs []string) int {

	minWordLength := math.MaxInt64

	for _, word := range strs {
		if len(word) < minWordLength {
			minWordLength = len(word)
		}
	}

	return minWordLength
}

func letterIsGeneralByIndex(strs []string, index int) bool {

	letter := strs[0][index]
	letterIsGeneralByIndex := true
	for _, word := range strs[1:] {
		if word[index] != letter {
			letterIsGeneralByIndex = false
			break
		}
	}

	return letterIsGeneralByIndex
}

// GetLongestCommonPrefix takes longest common prefix
func GetLongestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	sort.Strings(strs)

	minWordLength := getMinWordLength(strs)
	lastGeneralLetterIndex := -1
	for ind := 0; ind <= minWordLength-1; ind++ {
		if !letterIsGeneralByIndex(strs, ind) {
			break
		}
		lastGeneralLetterIndex = ind
	}

	if lastGeneralLetterIndex == -1 {
		return ""
	}

	longestCommonPrefix := strs[0][:lastGeneralLetterIndex+1]

	return longestCommonPrefix
}
