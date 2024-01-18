package main

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var r = new(ListNode)
	result := r
	var t1 int
	var t2 int
	var tt int
	flag := 0
	for ;l1 != nil || l2 !=nil; {
		if l1 == nil {
			t1 = 0
		} else {
			t1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			t2 = 0
		} else {
			t2 = l2.Val
			l2 = l2.Next
		}
		fmt.Println(t1, t2, tt)
		if flag > 0{
			r.Val = (t1 + t2 + tt)%10

		} else {
			r.Val = (t1 + t2)%10
		}
		tt = (t1 + t2 + tt)/10
		if tt > 0 || l1 != nil || l2 != nil {
			r.Next = new(ListNode)
		}
		r = r.Next
		flag++
	}
	return result
}

func main() {
	var l1 = new(ListNode)
	l1.Val = 2
	l1.Next = new(ListNode)
	l1.Next.Val = 4
	l1.Next.Next = new(ListNode)
	l1.Next.Next.Val = 3
	l1.Next.Next.Next = nil
	var l2 = new(ListNode)
	l2.Val = 5
	l2.Next = new(ListNode)
	l2.Next.Val = 6
	l2.Next.Next = new(ListNode)
	l2.Next.Next.Val = 4
	l2.Next.Next.Next = nil
	r := addTwoNumbers(l1, l2)
	for ;r != nil; {
		fmt.Println(r.Val)
		r = r.Next
	}
}