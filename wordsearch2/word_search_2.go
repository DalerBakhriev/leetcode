package wordsearch2

func findFirstLetterIndices(board [][]byte, firstLetter byte) [][]int {

	firstLetterIndices := make([][]int, 0)
	for indVertical := range board {
		for indHorizontal := range board[0] {
			if board[indVertical][indHorizontal] == firstLetter {
				firstLetterIndices = append(firstLetterIndices, []int{indVertical, indHorizontal})
			}
		}
	}

	return firstLetterIndices
}

func backTrack(
	results *[]string,
	targetWord string,
	currWord string,
	board [][]byte,
	currVertIndex int,
	currHorIndex int,
	currTargetWordIndex int,
	usedIndices map[[2]int]int,
) {

	if currWord == targetWord {
		*results = append(*results, currWord)
		return
	}

	if len(*results) > 0 {
		return
	}

	currTargetLetter := targetWord[currTargetWordIndex]
	candidatesIndices := [][]int{
		{currVertIndex, currHorIndex + 1},
		{currVertIndex, currHorIndex - 1},
		{currVertIndex + 1, currHorIndex},
		{currVertIndex - 1, currHorIndex},
	}

	for _, indicesCandidates := range candidatesIndices {
		firstInd, secondInd := indicesCandidates[0], indicesCandidates[1]
		_, indicesAlreadyUsed := usedIndices[[2]int{firstInd, secondInd}]
		if firstInd < len(board) && secondInd < len(board[0]) && firstInd >= 0 && secondInd >= 0 && !indicesAlreadyUsed {
			if board[firstInd][secondInd] == currTargetLetter {
				usedIndices[[2]int{firstInd, secondInd}] = 1
				backTrack(results, targetWord, currWord+string(currTargetLetter), board, firstInd, secondInd, currTargetWordIndex+1, usedIndices)
				delete(usedIndices, [2]int{firstInd, secondInd})
			}
		}
	}
}

// WordExist ...
func WordExist(board [][]byte, word string) bool {

	firstLetterIndices := findFirstLetterIndices(board, word[0])
	if len(firstLetterIndices) == 0 {
		return false
	}

	if len(word) == 1 && len(firstLetterIndices) > 0 {
		return true
	}

	results := make([]string, 0)
	for _, indicesOfFirstLetter := range firstLetterIndices {
		vertIndex, horIndex := indicesOfFirstLetter[0], indicesOfFirstLetter[1]
		usedIndices := make(map[[2]int]int)
		usedIndices[[2]int{vertIndex, horIndex}] = 1
		backTrack(&results, word, string(word[0]), board, vertIndex, horIndex, 1, usedIndices)
	}

	return len(results) > 0
}

func findWords(board [][]byte, words []string) []string {

	foundWords := make([]string, 0)

	for _, word := range words {
		if WordExist(board, word) {
			foundWords = append(foundWords, word)
		}
	}

	return foundWords
}
