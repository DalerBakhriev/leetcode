package lettertilepossibilities

import "strings"

func getSubsets(allSubsets map[string]int, candidates []string, currIndex int) {

	if currIndex == len(candidates) {
		allSubsets[""] = 1
		return
	}

	getSubsets(allSubsets, candidates, currIndex+1)
	currElement := candidates[currIndex]

	newSubsets := make([]string, 0)
	for subset := range allSubsets {
		newSubset := strings.Join([]string{subset, currElement}, "")
		newSubsets = append(newSubsets, newSubset)
	}

	for _, newSubset := range newSubsets {
		if _, alreadySaved := allSubsets[newSubset]; !alreadySaved {
			allSubsets[newSubset] = 1
		}
	}

}

func backTrack(
	solutions map[string]int,
	candidates string,
	combination string,
	combinationLength int,
	startIndex int,
	freqMap map[string]int,
) {

	if startIndex == combinationLength {
		if _, alreadyuSaved := solutions[combination]; !alreadyuSaved {
			solutions[combination] = 1
		}
		return
	}

	for ind := 0; ind < len(candidates); ind++ {
		candidate := string(candidates[ind])
		if freqMap[candidate] > 0 {
			freqMap[candidate]--
			backTrack(
				solutions,
				candidates,
				strings.Join([]string{combination, candidate}, ""),
				combinationLength,
				startIndex+1,
				freqMap,
			)
			freqMap[candidate]++
		}

	}
}

func getUniqueValues(s string) (map[string]int, []string) {

	uniqueVals := make(map[string]int)
	for _, str := range s {
		if _, ok := uniqueVals[string(str)]; !ok {
			uniqueVals[string(str)] = 1
		} else {
			uniqueVals[string(str)]++
		}
	}

	uniqueArr := make([]string, 0)
	for val := range uniqueVals {
		uniqueArr = append(uniqueArr, val)
	}

	return uniqueVals, uniqueArr
}

// NumTilePossibilities ...
func NumTilePossibilities(tiles string) int {

	allSubsets := make(map[string]int)
	candidates := strings.Split(tiles, "")
	getSubsets(allSubsets, candidates, 0)
	delete(allSubsets, "")

	tilePossibilities := make([]string, 0)
	uniqueTilePossibilities := make(map[string]int)
	for subset := range allSubsets {
		if len(subset) == 1 {
			tilePossibilities = append(tilePossibilities, subset)
		} else {
			freqMap, _ := getUniqueValues(subset)
			backTrack(uniqueTilePossibilities, subset, "", len(subset), 0, freqMap)
		}
	}

	for tilePossibility := range uniqueTilePossibilities {
		tilePossibilities = append(tilePossibilities, tilePossibility)
	}
	return len(tilePossibilities)
}
