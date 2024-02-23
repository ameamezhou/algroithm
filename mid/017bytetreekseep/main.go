package main

import (
	"sort"
)

// https://leetcode.cn/problems/kth-largest-sum-in-a-binary-tree/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	var temp = []int{}
	var accumulate func(t *TreeNode, deep int)
	accumulate = func(t *TreeNode, deep int) {
		if len(temp) < deep {
			temp = append(temp, t.Val)
		} else {
			temp[deep-1] += t.Val
		}
		if t.Right != nil {
			accumulate(t.Right, deep + 1)
		}
		if t.Left != nil {
			accumulate(t.Left, deep + 1)
		}
	}
	accumulate(root, 1)
	sort.Ints(temp)

	if k > len(temp) {
		return -1
	}

	return int64(temp[len(temp)-k])
}

