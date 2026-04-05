// strs_map = {} 
//
//
//
func groupAnagrams(strs []string) [][]string {
	freqs := make(map[[26]int][]string)

	for _, s := range strs {
		var count [26]int

		for _, c := range s {
			count[c-'a']++
		}
		freqs[count] = append(freqs[count], s)
	}

	var result [][]string

	for _, group := range freqs {
		result = append(result, group)
	}

	return result
}
