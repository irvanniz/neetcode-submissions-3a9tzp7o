func topKFrequent(nums []int, k int) []int {
	if len(nums) == 1 {
		return nums
	}

	mostFreqMap := make(map[int]int)
	var lenResult int
	// var result []int
	fmt.Println("k value: ", k)
	for i, num := range nums {
		// fmt.Println("num: ", num)
		// fmt.Println("val: ", mostFreqMap[num])
		
		// // at least already registered in map (>1 but avoid re-append in next iteration)
		// if mostFreqMap[num] == 0 {
		// 	result = append(result, num)
		// }

		// fmt.Printf("output [%d]: %+v \n", i+1, result)

		mostFreqMap[num]++
		if lenResult < mostFreqMap[num] {
			lenResult = mostFreqMap[num]
		}
		fmt.Printf("lenResult[%d]: %d \n", i, lenResult)
		// fmt.Println("map: ", mostFreqMap)
	}

	if lenResult == 1 {
		return nums
	}

	fmt.Println("final_map: ", mostFreqMap)

	tempResult := make([][]int, lenResult)
	for k, v := range mostFreqMap {
		fmt.Printf("map[%d]: %v \n", k, v)
		fmt.Println("prev_res: ", tempResult)
		tempResult[v-1] = append(tempResult[v-1], k)
		fmt.Println("next_res: ", tempResult)
	}

	fmt.Println("temp_res: ", tempResult)

	// if len(tempResult) >= k && 
	// 	len(tempResult[k-1]) == k {
	// 		return tempResult[k-1] 
	// }

	// totalResult:= len(result)
	// return result[len(result)-k:]

	var result []int
	lenTempResult := len(tempResult)
	for i := 0; i < lenTempResult; i++{
		fmt.Printf("right_result[%d]: %+v \n", i, tempResult[lenTempResult-1-i])
		result = append(result, tempResult[lenTempResult-1-i]...)
		fmt.Printf("f_result[%d]: %+v \n", i, result)

		if len(result) >= k {
			break
		}
	}

	return result
}
