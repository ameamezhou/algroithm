package main

func sumIndicesWithKSetBits(nums []int, k int) int {
	result := 0
	for i := range nums {
		if countBit(i) == k {
			result += nums[i]
		}
	}
	return result
}

func countBit(num int)(count int){
	for num > 0 {
		count += num % 2
		num /= 2
	}
	return
}
