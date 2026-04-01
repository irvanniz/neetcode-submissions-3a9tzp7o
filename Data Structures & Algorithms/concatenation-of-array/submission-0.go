func getConcatenation(nums []int) []int {
    var ans []int
	for i:=0;i<2;i++{
		ans = append(ans, nums...)
	}

	return ans
}
