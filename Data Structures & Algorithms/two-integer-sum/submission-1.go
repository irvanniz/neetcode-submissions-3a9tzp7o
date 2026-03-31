func twoSum(nums []int, target int) []int {
    numsMap := make(map[int]int)
    for idx, v := range nums {
        numsMap[v] = idx
    }

    for idx, v := range nums {
        complementer := target - v
        complIdx, isExist := numsMap[complementer]
        if isExist {
            if idx == complIdx {
                continue
            }
            
            return []int{idx, complIdx}
        }
    }

    return nil
}
