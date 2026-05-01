func twoSum(numbers []int, target int) []int {
	i, j := 0, 1
	for i < j {
		total := numbers[i] + numbers[j]
		if total == target {
			return []int{i+1, j+1}
		}

		j++
		if j >= len(numbers) {
			j = i+2
			if j > len(numbers) {
				j = len(numbers)-1
			}
			
			i++
		}
	}

	return []int{}
}
