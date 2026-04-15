func productExceptSelf(nums []int) []int {
	totalNums := len(nums)

	var (
		prefix []int
		p int = 1
	)
	for i:=0; i<totalNums; i++ {
		if i == 0 {
			prefix = append(prefix, 1)	
			continue
		}

		p*=nums[i-1]
		prefix = append(prefix, p)
	}

	var (
		suffix = make([]int, totalNums)
		s int = 1
	)
	
	for i:=(totalNums-1); i>=0; i-- {
		if i == 0 {
			suffix[totalNums-1] = 1
			continue
		}

		s*=nums[i]
		suffix[i-1] = s
	}

	var result []int
	for i, v := range prefix {
		result = append(result, v*suffix[i])
	}

	fmt.Printf("preffix: %+v | suffix: %+v | result: %+v\n", prefix, suffix, result)
	return result
}
