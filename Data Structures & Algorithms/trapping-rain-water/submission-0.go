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
			fmt.Printf("\ni: %d (%d)", i, height[i])
			maxLeft = max(maxLeft, height[i])
			totalWater += maxLeft - height[i]
			fmt.Printf(" | maxLeft: %d -> totalWater: %d", maxLeft, totalWater)

			i++
			continue
		}

		fmt.Printf("\nj: %d (%d)", j, height[j])
		maxRight = max(maxRight, height[j])
		totalWater += maxRight - height[j]
		fmt.Printf(" | maxRight: %d -> totalWater: %d", maxRight, totalWater)

		j--
	}

	return totalWater
}
