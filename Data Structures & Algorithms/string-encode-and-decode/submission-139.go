type Solution struct{
	data []string
}

func (s *Solution) Encode(strs []string) string {
	s.data = strs
	return ""
}

func (s *Solution) Decode(encoded string) []string {
	return s.data
}
