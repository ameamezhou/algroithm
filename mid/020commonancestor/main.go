package _20commonancestor

import "fmt"

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/description/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type cs struct {
	t	*TreeNode
	a 	int
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var result *TreeNode
	var combineSearch = []cs{}
	combineSearch = append(combineSearch, cs{t: root, a: -1})
	qSlice := []*TreeNode{}
	pSlice := []*TreeNode{}
	var stack = []*TreeNode{root}
	flag := 0
	for len(stack) > 0 {
		temp := stack[0]
		if temp.Left != nil {
			combineSearch = append(combineSearch, cs{t: temp.Left, a: flag})
			stack = append(stack, temp.Left)
		}
		if temp.Right != nil {
			combineSearch = append(combineSearch, cs{t: temp.Right, a: flag})
			stack = append(stack, temp.Right)
		}
		flag++
		stack = stack[1:]
	}
	var searchAncestor func(index int, tag string)
	searchAncestor = func(index int, tag string) {
		if tag == "q" {
			qSlice = append(qSlice, combineSearch[index].t)
		} else {
			pSlice = append(pSlice, combineSearch[index].t)
		}
		if combineSearch[index].a != -1 {
			searchAncestor(combineSearch[index].a, tag)
		}
	}

	for i, k := range combineSearch {
		if k.t.Val == p.Val {
			searchAncestor(i, "p")
		} else if k.t.Val == q.Val {
			searchAncestor(i, "q")
		}
	}
	fmt.Println(qSlice)
	fmt.Println(pSlice)
	iQ := 0
	iP := 0
	nQ := len(qSlice)
	nP := len(pSlice)
	if nQ > nP {
		iQ = nQ - nP
	} else {
		iP = nP - nQ
	}
	for iQ < nQ && iP < nP {
		if qSlice[iQ] == pSlice[iP] {
			break
		}
		iQ++
		iP++
	}
	result = qSlice[iQ]
	fmt.Println(result, result.Val)
	return result
}