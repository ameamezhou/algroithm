package main

// https://leetcode.cn/problems/make-costs-of-paths-equal-in-a-binary-tree/description/

// 题目很简单 根节点 -> 比较两个叶子节点的差值，补差值

func minIncrements(n int, cost []int) int {
	var result = 0
	var countFunc func(t int)
	countFunc = func(t int) {
		if t < 1 {
			return
		}
		result += abs(cost[t] - cost[t+1])
		cost[t/2] += max(cost[t], cost[t+1])
		countFunc(t-2)
	}
	countFunc(n-2)
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main(){
	tempSlice := []int{764,1460,2664,764,2725,4556,5305,8829,5064,5929,7660,6321,4830,7055,3761}
	minIncrements(15, tempSlice)
}