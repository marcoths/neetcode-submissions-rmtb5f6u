
func isValid(s string) bool {
	stack := make([]rune,0)
	symbols := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, c := range s {
		if open, ok := symbols[c]; ok {
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if top != open {
					return false
				}
			} else {
				return false
			}
		} else {
			stack = append(stack, c)
		}
	}

	return len(stack) == 0
    
}
