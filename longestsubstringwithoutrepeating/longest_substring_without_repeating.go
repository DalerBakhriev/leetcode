package longestsubstringwithoutrepeating

// LengthOfLongestSubstring returns length of longest substring without rrepeating
func LengthOfLongestSubstring(s string) int {

	var subStringLengths []int
	for ind := 0; ind <= len(s)-1; ind++ {
		subStringLengths = append(subStringLengths, LengthOfLongestSubstringFromSlice(s, ind))
	}

	var maxSubStringLength int
	for _, length := range subStringLengths {
		if length > maxSubStringLength {
			maxSubStringLength = length
		}
	}

	return maxSubStringLength
}

// LengthOfLongestSubstringFromSlice returns length of longest substring withour repeating starting with some index
func LengthOfLongestSubstringFromSlice(s string, startIndex int) int {

	var lengthOfLongestSubstring int
	lettersFromString := make(map[string]int)

	var currLongestSubstring int
	for _, letter := range s[startIndex:] {

		if _, ok := lettersFromString[string(letter)]; ok {
			lettersFromString = map[string]int{string(letter): 1}

			if currLongestSubstring > lengthOfLongestSubstring {
				lengthOfLongestSubstring = currLongestSubstring
			}

			currLongestSubstring = 1
			continue
		}

		currLongestSubstring++
		lettersFromString[string(letter)] = 1
	}

	if currLongestSubstring > lengthOfLongestSubstring {
		lengthOfLongestSubstring = currLongestSubstring
	}

	return lengthOfLongestSubstring
}

// LengthOfLongestSubstringFast fast way to find longest substring withourt repeat
func LengthOfLongestSubstringFast(s string) int {

	var lengthOfLongestSubstring int
	var currMaxLen int

	lettersLastIndex := make(map[rune]int)

	for ind, letter := range s {

		if lastIndex, ok := lettersLastIndex[letter]; ok && ind-currMaxLen <= lastIndex {

			if currMaxLen > lengthOfLongestSubstring {
				lengthOfLongestSubstring = currMaxLen
			}
			currMaxLen = ind - lastIndex
			lettersLastIndex[letter] = ind
			continue
		}

		currMaxLen++
		lettersLastIndex[letter] = ind
	}

	if currMaxLen > lengthOfLongestSubstring {
		lengthOfLongestSubstring = currMaxLen
	}

	return lengthOfLongestSubstring
}
