package main

import "fmt"

// https://leetcode.cn/problems/jump-game-vi/description/

func maxResult(nums []int, k int) int {
	var result = 0
	l := len(nums) - 1
	var tempFunc func(start, tempData int)
	tempFunc = func(start, tempData int) {
		tempData += nums[start]
		if start == l {
			if tempData > result {
				result = tempData
			}
		}
		for i := 1; i <= k; i++ {
			if start + i <= l {
				tempFunc(start + i, tempData)
			}
		}
	}
	tempFunc(0, 0)
	return result
}


func maxResult1(nums []int, k int) int {
	dp := make([]int, len(nums))
	if k > len(nums) {
		k = len(nums)
	}
	dp[0] = nums[0]
	preMax := dp[0]
	for i := 1; i < len(nums); i++ {
		if i <= k {
			dp[i] = preMax + nums[i]
			if dp[i] > preMax {
				preMax = dp[i]
			}
		} else {
			if dp[i-k-1] == preMax {
				preMax = max(dp[i-k : i]...)
			}
			dp[i] = preMax + nums[i]
			if dp[i] > preMax {
				preMax = dp[i]
			}
		}
	}
	return dp[len(nums)-1]
}

func max(nums ...int) (res int) {
	res = nums[0]
	for _, num := range nums {
		if num > res {
			res = num
		}
	}
	return
}



func maxResult2(nums []int, k int) int {
	l := len(nums)
	if k > l {
		k = l
	}
	tempValue := make([]int, l)
	tempValue[0] = nums[0]
	preMax := tempValue[0]
	for i := 1; i < l; i++ {
		if i <= k {
			tempValue[i] = preMax + nums[i]
			if tempValue[i] > preMax {
				preMax = tempValue[i]
			}
		} else {
			if tempValue[i - k - 1] == preMax {
				preMax = findMax(tempValue[i-k:i]...)
			}
			tempValue[i] = preMax + nums[i]
			if tempValue[i] > preMax {
				preMax = tempValue[i]
			}
		}
	}
	return tempValue[l - 1]
}

func findMax(nums ...int)(res int) {
	res = nums[0]
	for _, num := range nums {
		if num > res {
			res = num
		}
	}
	return
}

func main(){
	nums := []int{1, -5, -23, -30, 6}
	fmt.Println(maxResult2(nums, 2))
}