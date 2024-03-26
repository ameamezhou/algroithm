package _05maxPathSum

// https://leetcode.cn/problems/binary-tree-maximum-path-sum/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val int
    Left *TreeNode
    Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	var result = 0
	flag := 0
	var countPath func(t *TreeNode) int
	countPath = func(t *TreeNode) int {
		if t == nil {
			return 0
		}

		left := max(countPath(t.Left), 0)
		right := max(countPath(t.Right), 0)
		tempR := t.Val + left + right
		if flag == 0 {
			result = tempR
			flag = 1
		} else {
			result = max(tempR, result)
		}
		return t.Val + max(left, right)
	}
	countPath(root)
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}