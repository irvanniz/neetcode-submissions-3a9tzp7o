func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}

	s = strings.ToLower(s)

	var (
		s1 string
	)
	for _, v := range s {
		r := rune(v)
		if !isAlphanumeric(r) {
			continue
		}

		s1+=string(v)
	}

	i := 0
	j := len(s1) - 1
	for i < j {
		if s1[i] != s1[j] {
			return false
		}

		i++
		j--
	}

	return true
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