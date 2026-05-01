func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		total := numbers[i] + numbers[j]
		if total == target {
			return []int{i+1, j+1}
		}
		
		if total > target {
			j--
			continue
		}
		if total < target {
			i++
		}
	}

	return []int{}
}
