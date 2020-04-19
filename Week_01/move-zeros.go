func moveZeroes(nums []int) {
	j:=0
    for i,num :=range nums{
        if num!=0{
            nums[j],nums[i] = nums[i],nums[j]
            j++
        }
    }
}