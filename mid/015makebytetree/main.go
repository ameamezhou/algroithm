package main

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

func buildTree(preorder []int, inorder []int) *TreeNode {
	var result = new(TreeNode)

	var makeFunc func(t *TreeNode, pSlice []int, mSlice []int)
	makeFunc = func(t *TreeNode, pSlice []int, mSlice []int) {
		if len(pSlice) < 1 {
			return
		}
		t.Val = pSlice[0]
		k := 0
		n := len(mSlice)
		for ;k < n;k++ {
			if mSlice[k] == pSlice[0] {
				break
			}
		}
		lmSlice := mSlice[0:k]
		rmSlice := mSlice[k+1:]
		lpSlice := pSlice[1:k+1]
		rpSlice := pSlice[k+1:]
		if len(lmSlice) > 0 || len(lpSlice) > 0 {
			t.Left = new(TreeNode)
			makeFunc(t.Left, lpSlice, lmSlice)
		}
		if len(rmSlice) > 0 || len(rpSlice) > 0 {
			t.Right = new(TreeNode)
			makeFunc(t.Right, rpSlice, rmSlice)
		}
	}
	makeFunc(result, preorder, inorder)

	return result
}