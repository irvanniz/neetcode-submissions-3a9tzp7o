func topKFrequent(nums []int, k int) []int {
	if len(nums) == 1 {
		return nums
	}

	mostFreqMap := make(map[int]int)
	for _, num := range nums {
		mostFreqMap[num]++
	}

	var result resultData
	for k, v := range mostFreqMap {
		// if v <= 1 {
		// 	continue
		// }

		result = append(result, numData{
			num: k,
			total: v,
		})
		// if len(result) >= k {
		// 	break
		// }
	}

	sort.Slice(result, func(i, j int) bool{
		return result[i].total > result[j].total
	})

	// return result
	return result.output(k)
}

type resultData []numData

type numData struct {
	num int
	total int
}

func (r *resultData) output(limit int) []int {
	var result []int
	for _, v := range *r {
		result = append(result, v.num)

		if len(result) >= limit {
			return result
		}
	}

	return result
}
