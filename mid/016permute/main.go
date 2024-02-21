package main

import "sort"

// https://leetcode.cn/problems/permutations/

// 回溯

func permute(nums []int) [][]int {
	var result = [][]int{}
	n := len(nums)
	flag := make([]bool, 21)
	temp := make([]int, n)
	var backTrace func(index, count int)
	backTrace = func(index, count int) {
		flag[nums[index]+10] = true
		temp[count - 1] = nums[index]
		if count == n {
			var tempT = make([]int, n)
			copy(tempT, temp)
			result = append(result, tempT)
		} else {
			for k := range nums {
				if !flag[k] {
					backTrace(k, count + 1)
				}
			}
		}
		flag[nums[index]+10] = false
	}
	for k := range nums {
		backTrace(k, 1)
	}
	return result
}

// dfs
func permute2(nums []int) [][]int {

}

func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	perm := []int{}
	vis := make([]bool, n)
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), perm...))
			return
		}
		for i, v := range nums {
			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
				continue
			}
			perm = append(perm, v)
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	backtrack(0)
	return
}
