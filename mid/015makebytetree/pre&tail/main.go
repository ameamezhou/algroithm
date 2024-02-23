package main

// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-postorder-traversal/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	n := len(preorder)
	if n < 1 {
		return nil
	}
	var t = new(TreeNode)
	t.Val = preorder[0]
	if n == 1 {
		return t
	}
	i := 0
	t.Val = preorder[0] //根植
	leftVal := preorder[1] //左子树根值
	for ; i < n; i++ {
		if leftVal == postorder[i] {
			break
		}
	}
	i++
	t.Left = constructFromPrePost(preorder[1:1+i], postorder[:i])
	t.Right = constructFromPrePost(preorder[1+i:], postorder[i:n-1])

	return t
}