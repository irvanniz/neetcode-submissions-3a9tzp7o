func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{strs}
	}
	
	anagramMap := make(map[string][]string)
	
	for _, v := range strs {
		group := anagramGroup(v)

		anagramMap[group] = append(anagramMap[group], v)
	}

	var result [][]string
	for _, v := range anagramMap {
		result = append(result, v)
	}

	return result
}

func anagramGroup(str string) string {
	letters:=[]rune(str)

	sort.Slice(letters, func(i, j int) bool {
		return letters[i] < letters[j]
	})

	return string(letters)
}