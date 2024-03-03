package _03klinkreverse

// https://leetcode.cn/problems/reverse-nodes-in-k-group/description/

type ListNode struct {
    Val int
    Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k < 2 {
		return head
	}
	hair := &ListNode{ Next: head}
	pre := hair

	for head != nil {
		tail := pre
		for i:=0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		next := tail.Next
		head, tail = reverse(head, tail)
		pre.Next = head
		tail.Next = next
		pre = tail
		head = tail.Next
	}
	return hair.Next
}

func reverse(head, tail *ListNode)(*ListNode, *ListNode){
	h := head
	end := tail.Next
	for end != tail {
		next := h.Next
		h.Next = end
		end = h
		h = next
	}
	return tail, head
}