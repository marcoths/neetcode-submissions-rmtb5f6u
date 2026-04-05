import "maps"

func isAnagram(s string, t string) bool {
	sm := make(map[rune]int, len(s))
	tm := make(map[rune]int, len(t))

	for _, char := range s {
		sm[char]++
	}

	for _, char := range t {
		tm[char]++
	}

	return maps.Equal(sm, tm)

}
