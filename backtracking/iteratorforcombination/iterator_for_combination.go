package iteratorforcombination

import (
	"sort"
	"strings"
)

type CombinationIterator struct {
	Combinations []string
	CurrIndex    int
}

func backTrack(
	combinations map[string]int,
	characters string,
	combination string,
	combinationLength int,
	startIndex int,
) {

	if len(combination) > combinationLength {
		return
	}

	if len(combination) == combinationLength {
		stringsArr := strings.Split(combination, "")
		sort.Strings(stringsArr)
		sortedCombination := strings.Join(stringsArr, "")
		if _, ok := combinations[sortedCombination]; !ok {
			combinations[sortedCombination] = 1
		}
		return
	}

	for ind := startIndex; ind < len(characters); ind++ {
		if !strings.Contains(combination, string(characters[ind])) {
			backTrack(
				combinations,
				characters,
				strings.Join([]string{combination, string(characters[ind])}, ""),
				combinationLength,
				ind,
			)
		}
	}

}

// Constructor ...
func Constructor(characters string, combinationLength int) CombinationIterator {

	combinations := make(map[string]int, 0)
	backTrack(combinations, characters, "", combinationLength, 0)

	allCombinations := make([]string, 0)
	for combination := range combinations {
		allCombinations = append(allCombinations, combination)
	}

	sort.Strings(allCombinations)
	return CombinationIterator{
		Combinations: allCombinations,
		CurrIndex:    -1,
	}
}

// Next ...
func (combIterator *CombinationIterator) Next() string {
	combIterator.CurrIndex++

	return combIterator.Combinations[combIterator.CurrIndex]
}

// HasNext ...
func (combIterator *CombinationIterator) HasNext() bool {
	if combIterator.CurrIndex+1 < len(combIterator.Combinations) {
		return true
	}

	return false
}

/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
