package addandsearchword

type WordDictionary struct {
	wordsDict map[string]int
}

// Constructor Initialize your data structure here.
func Constructor() WordDictionary {
	wordsDict := make(map[string]int)

	return WordDictionary{wordsDict: wordsDict}
}

// AddWord Adds a word into the data structure
func (wD *WordDictionary) AddWord(word string) {
	if _, ok := wD.wordsDict[word]; !ok {
		wD.wordsDict[word] = 1
	}
}

func isMatch(s string, p string) bool {

	if len(p) == 0 {
		return len(s) == 0
	}

	currPairMatched := (len(s) > 0) && (p[0] == s[0] || p[0] == '.')

	return currPairMatched && isMatch(s[1:], p[1:])
}

// Search returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter.
func (wD *WordDictionary) Search(word string) bool {
	wordIsFound := false
	for wordFromDict := range wD.wordsDict {
		if isMatch(wordFromDict, word) {
			wordIsFound = true
			break
		}
	}

	return wordIsFound
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
