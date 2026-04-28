func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numsMap := make(map[int]bool)

	for _, v := range nums {
		numsMap[v] = true
	}

	var firstSeqNum []int
	for _, v := range nums {
		if !numsMap[v-1] {
			firstSeqNum = append(firstSeqNum, v)
		}
	}

	maxLength := 1
	for _, v := range firstSeqNum {
		tempMaxLength := 1
		i := v+1
		for numsMap[i]{
			i++
			tempMaxLength++
		}

		if tempMaxLength > maxLength {
			maxLength = tempMaxLength
		}
	}

	return maxLength
}
