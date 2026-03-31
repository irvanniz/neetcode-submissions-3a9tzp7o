func topKFrequent(nums []int, k int) []int {
	if len(nums) == 1 {
		return nums
	}

	mostFreqMap := make(map[int]int)
	var lenResult int
	
	for _, num := range nums {
		mostFreqMap[num]++
		if lenResult < mostFreqMap[num] {
			lenResult = mostFreqMap[num]
		}
	}

	tempResult := make([][]int, lenResult)
	for k, v := range mostFreqMap {
		tempResult[v-1] = append(tempResult[v-1], k)
	}

	var result []int
	lenTempResult := len(tempResult)
	for i := 0; i < lenTempResult; i++{
		result = append(result, tempResult[lenTempResult-1-i]...)

		if len(result) >= k {
			break
		}
	}

	return result
}
