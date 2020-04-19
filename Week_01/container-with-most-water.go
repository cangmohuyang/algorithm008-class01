func maxArea(height []int) int {
    i,j := 0,len(height)-1
    max :=0
    for i!=j {
         wide:=j-i;
         h:=0
         if height[i]>height[j]{
            h = height[j]
            j--
         }else{
            h = height[i]
            i++
         }
         area := h*wide
         if max <  area{
            max = area
         }
    }
    return max 
}