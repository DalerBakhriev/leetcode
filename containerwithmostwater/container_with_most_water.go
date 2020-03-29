package containerwithmostwater

func min(a, b int) int {

	if a < b {
		return a
	}

	return b
}

// MaxArea ...
func MaxArea(heights []int) int {

	var maxArea int
	for i := 0; i < len(heights); i++ {
		for j := 0; j < i; j++ {
			height := min(heights[i], heights[j])
			area := (i - j) * height
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}
