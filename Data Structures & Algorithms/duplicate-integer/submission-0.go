func hasDuplicate(nums []int) bool {
    numsMap := make(map[int]bool, len(nums))
    for _, v := range nums {
        _, isExist := numsMap[v]
        if isExist {
            return true
        }

        numsMap[v] = true
    }

    return false
}
