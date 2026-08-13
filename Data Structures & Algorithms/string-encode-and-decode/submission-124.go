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
	if encoded == "" {
		return nil
	}
	
	var result []string
	for i:=0; i < len(encoded); i++ {
		j := i
		var lenItemStr string
		for string(encoded[j]) != s.key {
			if isLenItem(rune(encoded[j])) {
				lenItemStr+=string(encoded[j])
			}else{
				break
			}
			j++
		}
		
		lenItem, _ := strconv.Atoi(string(lenItemStr))
		item := string(encoded[j+1:j+1+lenItem])
		result = append(result, item)
		i = j+lenItem
	}

	return result
}

func isLenItem(s rune) bool{
	return s >= '0' && s <= '9'
}
