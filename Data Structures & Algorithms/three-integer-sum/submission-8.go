import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	res := [][]int{}
	for i, n := range nums {
		left, right := i + 1, len(nums)-1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for left < right {
			
			cur := n + nums[left] + nums[right]

			if cur == 0 {
				res =append(res, []int{n, nums[left], nums[right]})
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			} else if cur < 0 {
				left++
			} else {
				right--
			}
			
			if nums[left] == nums[left-1] {
				continue
			}
		}
	}
	return res

}

//	8 
//. l       n
// [1,2,3,4,5]