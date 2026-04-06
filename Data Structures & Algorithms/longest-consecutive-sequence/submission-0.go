//
func longestConsecutive(nums []int) int {

	numsSet := make(map[int]bool)
	longest := 0

	for _, n := range nums {
		if _, ok := numsSet[n]; !ok {
			numsSet[n] = true
		}
	}

	for _, n := range nums {
		prev := n - 1
		if _, ok := numsSet[prev]; !ok {
			currSeq := 1
			for {
				if _, ok := numsSet[n+currSeq]; ok {
					currSeq++
				} else {
					break
				}
			}
			if currSeq > longest {
				longest = currSeq
			}
		}
	}
	return longest



	// {0,1,2,3,4,5,6}
}
