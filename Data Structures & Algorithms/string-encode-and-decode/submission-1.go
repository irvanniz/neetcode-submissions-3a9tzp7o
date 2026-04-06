type Solution struct{
}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 1 && strs[0] == "" {
		return ""
	}
	
	enc, err := json.Marshal(strs)
	if err != nil {
		fmt.Printf("[Encode] failed unmarshal input: %+v, err: %w \n", strs, err)
		return ""
	}

	return string(enc)
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "" {
		return []string{""}
	}

	var result []string
	err := json.Unmarshal([]byte(encoded), &result)
	if err != nil {
		fmt.Printf("[Decode] failed unmarshal input: %s, err: %w \n", encoded, err)
		return nil
	}

	return result
}
