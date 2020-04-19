#### 	学习笔记

​	本周完成了五个算法题，分别是移动零、爬楼梯、三数之和、环形链表、两个数之和、盛水最多的容器。

 1. 移动零。优化之后的算法使用的是，双指针。

     1. 指针`i`一直在移动

     2. 一直到遇到非零的数`i`，`j`位置的数字交换，j++

     3. 此处的交换比较巧妙

        ```go
        func moveZeroes(nums []int) {
        	j:=0
            for i,num :=range nums{
                if num!=0{
                    nums[j],nums[i] = nums[i],nums[j]
                    j++
                }
            }
        }
        ```

        正常的思维程序如下：

        ```go
        func moveZeroes(nums []int) {
        	j:=0
            for i,num :=range nums{
                if num!=0{
                    nums[j]=nums[i]
                    //填充补0
                    if i!=j{
                        nums[i] = 0
                    }
                    j+=1
                }
            }
        }
        ```

    2. 爬楼梯。解题思路，递归->递归，循环展开->储存的中间数据的递归

       ```go
       var m = make(map[int]int)
       func climbStairs(n int) int {
       	return climb_Stairs(n)
       }
       
       func climb_Stairs(n int) int {
       	if n == 1 {
       		return 1
       	}
       	if n == 2 {
       		return 2
       	}
       	i, j := 0, 0
       	if v, ok := m[n-1]; ok {
       		i = v
       	} else {
       		i = climb_Stairs(n - 1)
       		m[n-1] = i
       	}
       	if v, ok := m[n-2]; ok {
       		j = v
       	} else {
       		j = climb_Stairs(n - 2)
       		m[n-2] = j
       	}
       	return i + j
       }
       ```

    3. 两个数字之和为0。通过哈希表

       ```go
       func twoSum(nums []int, target int) []int {
           m:=make(map[int]int)
       
           for i,num :=range nums{
               if v,ok:=m[target-num];ok{
                   return []int{i,v}
               }else{
                   m[num]=i
               }
           }
           return nil
       }
       ```

4. 三数之和。暴力求解->两层循环+两数之和->sort之后，双指针移动，指针分别从开始和结尾的位置开始，其中，为了解决重复问题，需要判断`nums[k]==nums[k-1]`

   ```go
   func threeSum(nums []int) [][]int {
   	if nums == nil && len(nums) < 3 {
   		return nil
   	}
   	res := make([][]int, 0)
   	sort.Ints(nums)
   	for k, _ := range nums {
   		if nums[k] > 0 {
   			return res
   		}
   		i := k + 1
   		j := len(nums) - 1
   		if k > 0 && nums[k] == nums[k-1] {
   			continue
   		}
   		for i < j {
   			sum := nums[k] + nums[i] + nums[j]
   			if sum < 0 {
   				i++
   			} else if sum > 0 {
   				j--
   			} else {
   				res = append(res, []int{nums[k], nums[i], nums[j]})
   				for i<j&&nums[i] ==nums[i+1]{
                       i++
                   }
                   for i<j&&nums[j] ==nums[j-1]{
                       j--
                   }
                   i++
                   j--
   			}
   		}
   	}
   	return res
   }
   ```

   5. 盛最多水的容器。暴力求解->双指针，从开始和结尾的位置开始，其中那个高度小，就移动此指针。

      ```
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
      ```

      

