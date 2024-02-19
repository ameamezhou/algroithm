package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */


type Node struct {
    Val int
    Children []*Node
}



func postorder(root *Node) []int {
	var result = []int{}

	var orderFunc func(n *Node)
	orderFunc = func(n *Node) {
		for _, v := range n.Children {
			orderFunc(v)
		}
		result = append(result, n.Val)
	}
	orderFunc(root)
	return result
}
