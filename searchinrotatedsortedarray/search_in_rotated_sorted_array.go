package searchinrotatedsortedarray

func search(nums []int, target int) int {
	index := -1
	for ind, num := range nums {
		if num == target {
			index = ind
			break
		}
	}
	return index
}
