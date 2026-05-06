func maxArea(heights []int) int {
	var (
		maxAmount int
		i int
		j = len(heights)-1
	)

	for i<j {
		var (
			result int
			barRange = (j+1)-(i+1)
		)
		if heights[i] < heights[j]{
			result = heights[i] * barRange
			i++
		}else{
			result = heights[j] * barRange
			j--
		}

		if result > maxAmount {
			maxAmount = result
		}
	}

	return maxAmount
}
