package _83movezero

// https://leetcode.cn/problems/move-zeroes/description/?envType=study-plan-v2&envId=top-100-liked

func moveZeroes(nums []int)  {
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}