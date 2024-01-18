package main

import "fmt"

/*
https://leetcode.cn/problems/jump-game-ii/

给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。

每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:

0 <= j <= nums[i]
i + j < n
返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。



示例 1:

输入: nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
示例 2:

输入: nums = [2,3,0,1,4]
输出: 2
*/

func searchPath(nums []int) int {
	head := 0
	temp_rear := len(nums) - 1
	rear := temp_rear
	var result int
	for head < rear {
		if head + nums[head] >= rear {
			result = head
			break
		}
		if temp_rear + nums[temp_rear] >= rear {
			result = temp_rear
		}
		temp_rear--
		head++
	}
	return result
}

func jump(nums []int) int {
	count := 0

	for index := len(nums) - 1 ; index > 0; index = searchPath(nums) {
		nums = nums[0:index+1]
		count++
	}
	return count
}

func main(){
	nums := []int{1, 2, 3}
	fmt.Println(jump(nums))
}

func jump2(nums []int) int {
	var (
		count = 0
		next = 0
		maxPos = 0
	)
	for i := 0; i < len(nums) - 1; i++ {
		if i + nums[i] > maxPos {
			maxPos = i + nums[i]
		}
		if i == next {
			count++
			next = maxPos
		}
	}
	return count
}