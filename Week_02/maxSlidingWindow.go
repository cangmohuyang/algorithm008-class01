func maxSlidingWindow(nums []int, k int) []int {
    res:=make([]int,0)
    if len(nums) == 0{
        return nil
    }
    for i:=0;i<=len(nums)-k;i++{
       res= append(res,max(nums[i:k+i]))
    }
    return res
}
func max(nums[]int)int{
    max:=nums[0]
    for _,v:=range nums{
        if max < v{
            max =v
        }
    }
    return max
}