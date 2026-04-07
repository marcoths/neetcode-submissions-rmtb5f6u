func minEatingSpeed(piles []int, h int) int {

	end := 0

	for _, n := range piles {
		if n > end {
			end = n
		}
	}
	res := end
	start := 1

	for start <= end {
		mid := 	start + (end - start) / 2
		hours := 0
		for _, p := range piles {
			hours += int(math.Ceil(float64(p) / float64(mid)))
		}
		if hours <= h {
			res = min(res, mid)
			end = mid - 1
		} else {
			start = mid + 1 
		}
	}

	return res

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}