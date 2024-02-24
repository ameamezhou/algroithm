package main

import "sort"

// https://leetcode.cn/problems/closest-nodes-queries-in-a-binary-search-tree/description/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type qTree struct {
	node 	*TreeNode
	storey 	int
}

func closestNodes1(root *TreeNode, queries []int) [][]int {
	n := len(queries)
	var flag = make([]bool, n)
	var result = make([][]int, n)
	for i := range result {
		result[i] = []int{-1, -1}
	}
	queue := []*TreeNode{}
	queue = append(queue, root)

	var tempT *TreeNode
	for len(queue) > 0 {
		tempT = queue[0]
		for k, v := range queries {
			if !flag[k] {
				if v == tempT.Val {
					result[k] = []int{v, v}
					flag[k] = true
				} else {
					if tempT.Val < v {
						if tempT.Val > result[k][0] {
							result[k][0] = tempT.Val
						}
					} else if tempT.Val > v {
						if result[k][1] == -1 || tempT.Val < result[k][1] {
							result[k][1] = tempT.Val
						}
					}
				}
			}
		}
		if tempT.Left != nil {
			queue = append(queue, tempT.Left)
		}
		if tempT.Right != nil {
			queue = append(queue, tempT.Right)
		}
		queue = queue[1:]
	}
	return result
}

func closestNodes(root *TreeNode, queries []int) [][]int {
	arr := []int{}
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		arr = append(arr, root.Val)
		dfs(root.Right)
	}

	dfs(root)
	res := make([][]int, len(queries))
	for i, val := range queries {
		maxVal, minVal := -1, -1
		index := sort.SearchInts(arr, val)
		if index < len(arr) {
			maxVal = arr[index]
			if arr[index] == val {
				minVal = arr[index]
				res[i] = []int{minVal, maxVal}
				continue
			}
		}
		if index != 0 {
			minVal = arr[index - 1]
		}
		res[i] = []int{minVal, maxVal}
	}
	return res
}

func main(){

}
