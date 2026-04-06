type Solution struct{
	key string
}

func (s *Solution) Encode(strs []string) string {
	s.key = "#"
	
	var encoded string
	for _, v := range strs {
		delimiter := fmt.Sprintf("%d%s", len(v), s.key)

		encoded+=delimiter+v
	}

	return encoded
}

func (s *Solution) Decode(encoded string) []string {
	fmt.Printf("[Decode] enc: %s, len: %d \n\n", encoded, len(encoded))
	if encoded == "" {
		return nil
	}
	
	var result []string
	for i:=0; i < len(encoded); i++ {
		fmt.Printf("\n\n[Decode] i: %d", i)
		j := i
		var lenItemStr string
		for string(encoded[j]) != s.key {
			fmt.Printf(" - j: %d(%T) | encJ: %+v(%T) | isLenItem: %v \n", j,j, encoded[j],encoded[j], isLenItem(rune(encoded[j])))
			if isLenItem(rune(encoded[j])) {
				lenItemStr+=string(encoded[j])
			}else{
				break
			}
			j++
		}
		
		lenItem, _ := strconv.Atoi(string(lenItemStr))
		fmt.Printf("[Decode] lenItem: %d \n", lenItem)

		item := string(encoded[j+1:j+1+lenItem])
		result = append(result, item)

		// fmt.Printf("[Decode] v: %s \n", string(encoded[i]))
		// if string(encoded[i]) == s.key {
		// 	lenItem, _ := strconv.Atoi(string(encoded[i-1]))
		// 	firstWord := i+1
		// 	item := encoded[firstWord:firstWord+lenItem]

		// 	fmt.Println("len: ", lenItem)
			fmt.Println("item: ", item)
		// 	fmt.Printf("--------------\n")
		// 	result = append(result, item)
			i = j+lenItem
		// 	fmt.Println("total: ", len(encoded))
		// 	fmt.Println("i: ", i)
		// }
	}

	return result
}

func isLenItem(s rune) bool{
	fmt.Printf("[isLenItem] s: %+v (%T) | str :%s \n", s,s, string(s))
	// return s >= 48 && s <= 57
	return s >= '0' && s <= '9'
}
