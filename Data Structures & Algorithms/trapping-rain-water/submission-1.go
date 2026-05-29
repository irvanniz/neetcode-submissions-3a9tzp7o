func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	var (
		maxLeft, maxRight, totalWater int
		i int
		j = len(height)-1
	)

	for i<=j{
		if maxLeft < maxRight {
			maxLeft = max(maxLeft, height[i])
			totalWater += maxLeft - height[i]
			i++
			continue
		}
	
		maxRight = max(maxRight, height[j])
		totalWater += maxRight - height[j]
		j--
	}

	return totalWater
}
