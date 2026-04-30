func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}

	s = strings.ToLower(s)

	var s1 string
	for _, v := range s {
		fmt.Printf("v: %+v-%s(%T)\n", rune(v), string(v), v)
		if !isAlphanumeric(rune(v)) {
			continue
		}

		s1+=string(v)
	}

	s2 := reverse(s1)
	if s1 == s2 {
		return true
	}

	fmt.Printf("\ns: %s \ns1: %s \ns2: %s", s, s1, s2)
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

// this func below copas from google with keyword "golang reverse string"
func reverse(s string) string {
    // Convert string to a slice of runes
    runes := []rune(s)
    
    // Use two pointers to swap elements in place
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    
    return string(runes)
}