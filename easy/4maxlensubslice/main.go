package main



func alternatingSubarray(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	result := 0
	n := len(nums) - 1
	for i := 0; i < n; {
		temp := 0
		if nums[i+1] - nums[i] != 1 {
			i++
			continue
		}
		temp += 2
		i += 2
		for i <= n && nums[i] == nums[i-2] {
			i++
			temp++
		}
		if temp > result {
			result = temp
		}
		i--
	}
	if result != 0 {
		return result
	}
	return -1
}
