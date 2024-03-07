package _917

// https://leetcode.cn/problems/find-the-k-or-of-an-array/description/?envType=daily-question&envId=2024-03-06


func findKOr(nums []int, k int) int {
	ans := 0
	for i := 0; i < 31; i++ {
		cnt := 0
		for _, num := range nums {
			if (num >> i) & 1 == 1 {
				cnt++
			}
		}
		if cnt >= k {
			ans |= 1 << i
		}
	}
	return ans
}