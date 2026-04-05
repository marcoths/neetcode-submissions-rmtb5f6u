
func twoSum(nums []int, target int) []int {
    complements := make(map[int]int)

	for i, n := range nums {
		comp := target - n 

		if idx, ok := complements[comp]; ok {
			return []int{idx, i}			
		} 

		complements[n] = i
	}

	return []int{0, 0}
}
