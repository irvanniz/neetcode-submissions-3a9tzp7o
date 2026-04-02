type NumArray struct {
    data []int
}


func Constructor(nums []int) NumArray {
	return NumArray{
		data: nums,
	}
}


func (this *NumArray) SumRange(left int, right int) int {
	var result int
	for i:=left; i<=right; i++ {
		result+=this.data[i]
	}
	
    return result
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */