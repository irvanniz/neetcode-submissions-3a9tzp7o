import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	var (
		result [][]int
	)
	for i, v := range nums {
		target := v * -1
		j, k := i+1, len(nums)-1

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j < k {
			total := nums[j] + nums[k]
			if total == target {
				t := []int{v, nums[j], nums[k]}
				result = append(result, t)

				j++
				k--
				for j < k && nums[j] == nums[j-1] { j++ }
				for j < k && nums[k] == nums[k+1] { k-- }
				continue
			}

			if total > target {
				k--
				continue
			}
			if total < target {
				j++
			}
		}

	}

	return result
}
