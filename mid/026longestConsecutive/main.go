package _26longestConsecutive

// https://leetcode.cn/problems/longest-consecutive-sequence/?envType=study-plan-v2&envId=top-100-liked

func longestConsecutive(nums []int) int {
	tempCache :=  map[int]bool{}
	for _, v := range nums {
		tempCache[v] = true
	}
	result := 0
	for num := range tempCache{
		if !tempCache[num-1] {
			tempNum := num
			tempLen := 1
			for tempCache[tempNum+1] {
				tempNum++
				tempLen++
			}
			if tempLen > result {
				result = tempLen
			}
		}
	}
	return result
}