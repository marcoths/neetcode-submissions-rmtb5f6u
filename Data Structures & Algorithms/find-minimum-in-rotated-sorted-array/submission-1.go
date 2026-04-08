func findMin(nums []int) int {
	lo, hi := 0, len(nums) - 1

	min := -1

	for lo <= hi {
		mid := lo + (hi - lo) / 2
		if nums[mid] <= nums[len(nums)-1] {
			min = nums[mid]
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return min
}
