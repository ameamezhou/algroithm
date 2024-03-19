package _27dself

// https://leetcode.cn/problems/product-of-array-except-self/?envType=study-plan-v2&envId=top-100-liked

func productExceptSelf(nums []int) []int {
	flag := 0
	tempResult := 1
	n := len(nums)
	tempIndex := 0
	for k, v := range nums {
		if v == 0 {
			flag++
			if flag == 2 {
				break
			}
			tempIndex = k
			continue
		}
		tempResult *= v
	}
	var result = make([]int, n)
	if flag == 0 {
		for k, v := range nums {
			result[k] = tempResult/v
		}
	} else if flag == 1 {
		result[tempIndex] = tempResult
	}
	return result
}
