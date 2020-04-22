package maxlengthofconcatstrwithuniquechars

func backTrack(
	uniqueCharsStr map[string]int,
	candidates []string,
	currStr string,
	startIndex int,
) {

	if _, ok := uniqueCharsStr[currStr]; !ok && currStr != "" {
		if checkAllCharsAreUnique(currStr) {
			uniqueCharsStr[currStr] = len(currStr)
		}
	}

	if startIndex == len(candidates) {
		return
	}

	for _, shouldAppend := range []bool{true, false} {
		nextStr := currStr
		if shouldAppend {
			nextStr = currStr + candidates[startIndex]
		}
		backTrack(
			uniqueCharsStr,
			candidates,
			nextStr,
			startIndex+1,
		)
	}
}

func checkAllCharsAreUnique(str string) bool {

	chars := make(map[rune]int)

	allCharsAreUnique := true
	for _, char := range str {
		if _, ok := chars[char]; ok {
			allCharsAreUnique = false
			break
		}
		chars[char] = 1
	}

	return allCharsAreUnique
}

// MaxLength ...
func MaxLength(arr []string) int {

	if len(arr) == 1 {
		if checkAllCharsAreUnique(arr[0]) {
			return len(arr[0])
		}
		return 0
	}

	uniquesCharsStrings := make(map[string]int)
	backTrack(uniquesCharsStrings, arr, "", 0)

	currMaxLength := 0
	for _, length := range uniquesCharsStrings {
		if length > currMaxLength {
			currMaxLength = length
		}
	}

	return currMaxLength
}
