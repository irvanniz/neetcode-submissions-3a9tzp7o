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
	fmt.Println("str: "+str)
	count := make([]int, 26)  // a-z
    for _, ch := range str {
        count[ch-'a']++
    }
	fmt.Println("count: ",count)
    return fmt.Sprintf("%v", count)
}