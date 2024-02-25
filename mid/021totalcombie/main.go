package _21totalcombie

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/combination-sum/

func combinationSum(candidates []int, target int) [][]int {
	var result = [][]int{}
	var combineSum func(tempSlice []int)
	combineSum = func(tempSlice []int) {
		if add(tempSlice, target) == 1 {
			fmt.Println(tempSlice)
			result = append(result, tempSlice)
			return
		} else if add(tempSlice, target) == 2 {
			return
		}
		for _, v := range candidates {
			fmt.Println(1)
			tempSlice = append(tempSlice, v)
			combineSum(tempSlice)
		}
	}
	combineSum([]int{})
	return result
}

func add(tempS []int, target int) int {
	temp := 0
	for _, v := range tempS {
		temp += v
	}
	if target == temp {
		return 1
	} else if target > temp {
		return 2
	}
	fmt.Println(tempS)
	return 0
}

var (
	res [][]int
	path  []int
)
func combinationSum1(candidates []int, target int) [][]int {
	res, path = make([][]int, 0), make([]int, 0, len(candidates))
	sort.Ints(candidates)   // 排序，为剪枝做准备
	dfs(candidates, 0, target)
	return res
}

func dfs(candidates []int, start int, target int) {
	if target == 0 {   // target 不断减小，如果为0说明达到了目标值
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	for i := start; i < len(candidates); i++ {
		if candidates[i] > target {  // 剪枝，提前返回
			break
		}
		path = append(path, candidates[i])
		dfs(candidates, i, target - candidates[i])
		path = path[:len(path) - 1]
	}
}