package medianoftwosortedarrays

func getMedianIndexes(arraySize int) (int, int) {

	firstIndex := arraySize / 2
	secondIndex := firstIndex

	if arraySize%2 == 0 {
		secondIndex = firstIndex - 1
	}

	return firstIndex, secondIndex
}

// FindMedianSortedArrays finds median of arrays
func FindMedianSortedArrays(arrayOne, arrayTwo []int) float64 {

	var resultMergedArray []int
	var lastIndexArrayOne int
	var lastIndexArrayTwo int

	var firstMedianTerm, secondMedianTerm int
	var mergedArraySize = len(arrayOne) + len(arrayTwo)
	firstMedianIndex, secondMedianIndex := getMedianIndexes(mergedArraySize)

	for len(resultMergedArray) < mergedArraySize {

		if lastIndexArrayOne > len(arrayOne)-1 {
			resultMergedArray = append(resultMergedArray, arrayTwo[lastIndexArrayTwo:]...)
		} else if lastIndexArrayTwo > len(arrayTwo)-1 {
			resultMergedArray = append(resultMergedArray, arrayOne[lastIndexArrayOne:]...)
		}

		if len(resultMergedArray)-1 >= firstMedianIndex {
			firstMedianTerm = resultMergedArray[firstMedianIndex]
			secondMedianTerm = resultMergedArray[secondMedianIndex]
			break
		}

		if arrayOne[lastIndexArrayOne] < arrayTwo[lastIndexArrayTwo] {
			resultMergedArray = append(resultMergedArray, arrayOne[lastIndexArrayOne])
			lastIndexArrayOne++
		} else if arrayTwo[lastIndexArrayTwo] < arrayOne[lastIndexArrayOne] {
			resultMergedArray = append(resultMergedArray, arrayTwo[lastIndexArrayTwo])
			lastIndexArrayTwo++
		} else {
			resultMergedArray = append(resultMergedArray, arrayOne[lastIndexArrayOne], arrayTwo[lastIndexArrayTwo])
			lastIndexArrayOne++
			lastIndexArrayTwo++
		}

		if len(resultMergedArray)-1 >= firstMedianIndex {
			firstMedianTerm = resultMergedArray[firstMedianIndex]
			secondMedianTerm = resultMergedArray[secondMedianIndex]
			break
		}
	}

	return float64(firstMedianTerm+secondMedianTerm) / 2

}
