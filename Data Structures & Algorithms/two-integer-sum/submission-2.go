func twoSum(nums []int, target int) []int {
    numsMap := make(map[int]int)
    for idx, v := range nums {
        complementer := target - v
        complIdx, isExist := numsMap[complementer]
        if isExist {
            return []int{complIdx, idx}
        }

        numsMap[v] = idx
    }

    return nil
}
