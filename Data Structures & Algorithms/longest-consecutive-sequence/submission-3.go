func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var (
		numsMap = make(map[int]bool)
	)

	// build map
	for _, v := range nums {
		numsMap[v] = true
	}

	// find the first sequence candidate
	var firstSeqNum []int
	for _, v := range nums {
		if !numsMap[v-1] {
			firstSeqNum = append(firstSeqNum, v)
		}
	}

	fmt.Printf("firstSeqNum: %+v \n", firstSeqNum)

	// find the max length of consecutive sequence
	maxLength := 1
	for _, v := range firstSeqNum {
		tempMaxLength := 1
		for i:=1; i<len(nums); i++{
			fmt.Printf("[%d]v+i: %d \n", v, v+i)
			if !numsMap[v+i] {
				break
			}

			tempMaxLength++
		}

		if tempMaxLength > maxLength {
			maxLength = tempMaxLength
		}
	}

	// fmt.Printf("numsMap: %+v \nconsNumsMap: %+v \n", numsMap, consNumsMap)
	return maxLength
}
