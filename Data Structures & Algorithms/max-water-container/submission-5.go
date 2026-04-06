func maxArea(heights []int) int {
	best := 0
	left := 0
	right := len(heights)-1

	for left < right {
		area := (right - left) 	* min(heights[left], heights[right])
		best = max(best, area)

		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}

	return best

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
