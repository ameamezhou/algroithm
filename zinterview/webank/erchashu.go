package webank

import (
	"fmt"
	"math"
)

/*
给一颗二叉树，写程序找出所有从根节点到叶子节点路径中节点总和最大的路径。如果最大路径有多个，请全部输出。

输入:[5,4,8,11,null,9,4,7,2,null,null,5,1]
输出:[[5,4,11,7],[5,8,9,5]]
*/


/////////////////////////////////////////////

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) [][]int {
	var paths [][]int
	var maxSum int

	var dfs func(node *TreeNode, path []int, sum int)
	dfs = func(node *TreeNode, path []int, sum int) {
		if node == nil {
			return
		}

		path = append(path, node.Val)
		sum += node.Val

		if node.Left == nil && node.Right == nil {
			if sum > maxSum {
				maxSum = sum
				paths = [][]int{append([]int(nil), path...)}
			} else if sum == maxSum {
				paths = append(paths, append([]int(nil), path...))
			}
			return
		}

		dfs(node.Left, path, sum)
		dfs(node.Right, path, sum)
	}

	dfs(root, []int{}, 0)

	return paths
}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}

	paths := maxPathSum(root)
	fmt.Println(paths)
}

func maxPathSum1(nums []int) [][]int {
	var paths [][]int
	var maxSum int

	var dfs func(index int, path []int, sum int)
	dfs = func(index int, path []int, sum int) {
		if index >= len(nums) {
			return
		}

		path = append(path, nums[index])
		sum += nums[index]

		leftChild := 2*index + 1
		rightChild := 2*index + 2

		if leftChild >= len(nums) && rightChild >= len(nums) {
			if sum > maxSum {
				maxSum = sum
				paths = [][]int{append([]int(nil), path...)}
			} else if sum == maxSum {
				paths = append(paths, append([]int(nil), path...))
			}
			return
		}

		dfs(leftChild, path, sum)
		dfs(rightChild, path, sum)
	}

	dfs(0, []int{}, 0)

	return paths
}

func main1() {
	nums := []int{5, 4, 8, 11, math.MinInt32, 9, 4, 7, 2, math.MinInt32, math.MinInt32, 5, 1}
	paths := maxPathSum1(nums)
	fmt.Println(paths)
}
