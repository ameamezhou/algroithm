package main

// https://leetcode.cn/problems/range-sum-of-bst/

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	var result = 0
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		temp := queue[0]

		if temp.Val >= low && temp.Val <= high {
			result += temp.Val
		}

		if temp.Left != nil {
			queue = append(queue, temp.Left)
		}

		if temp.Right != nil {
			queue = append(queue, temp.Right)
		}

		queue = queue[1:]
	}
	return result
}