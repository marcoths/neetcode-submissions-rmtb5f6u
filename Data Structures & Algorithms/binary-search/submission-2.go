func search(nums []int, target int) int {
	lo, hi := 0, len(nums) - 1

	for lo <= hi {
		mid := lo + (hi - lo) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			hi--
		} else {
			lo++
		}
	}
	return - 1
}
