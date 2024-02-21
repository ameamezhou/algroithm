package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(inorder)
	if n < 1 {
		return nil
	}
	var result = new(TreeNode)
	tempVal := postorder[n-1]
	result.Val = tempVal
	var k = 0
	for ; k < n; k++ {
		if tempVal == inorder[k] {
			break
		}
	}
	result.Left = buildTree(inorder[0:k], postorder[0:k])
	result.Right = buildTree(inorder[k+1:], postorder[k:n-1])
	return result
}