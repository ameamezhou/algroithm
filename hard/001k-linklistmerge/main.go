package _01k_linklistmerge

// https://leetcode.cn/problems/merge-k-sorted-lists/description/

type ListNode struct {
    Val int
    Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var result *ListNode
	if len(lists) < 1 {
		return result
	}
	result = lists[0]
	for i:=1; i < len(lists); i++{
		result = merge(result, lists[i])
	}
	return result
}

func merge(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				cur.Next = l1
				l1 = l1.Next
			} else {
				cur.Next = l2
				l2 = l2.Next
			}
			cur = cur.Next
		} else if l1 != nil {
			cur.Next = l1
			break
		} else {
			cur.Next = l2
			break
		}
	}
	return head.Next
}