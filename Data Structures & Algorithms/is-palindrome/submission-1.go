func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}

	s = strings.ToLower(s)

	var (
		s1 string
		s2R []rune
	)
	for _, v := range s {
		r := rune(v)
		if !isAlphanumeric(r) {
			continue
		}

		s1+=string(v)
		s2R = append([]rune{r}, s2R...)
	}

	s2 := string(s2R)
	if s1 == s2 {
		return true
	}

	return false
}

func isAlphanumeric(v rune) bool{
	if v >= '0' && v <= '9' {
		return true
	}
	if v >= 'a' && v <= 'z' {
		return true
	}

	return false
}