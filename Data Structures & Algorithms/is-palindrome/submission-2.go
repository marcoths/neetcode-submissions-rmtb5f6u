func isPalindrome(s string) bool {
	left, right := 0, len(s) -1

	for left < right {
		for left < right && !isAlnum(rune(s[right])) {
			right -= 1
		}

		for left < right && !isAlnum(rune(s[left])) {
			left += 1
		}

		if strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])) {
			return false
		}

		left += 1
		right -= 1
	}
	return true
}

func isAlnum(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
