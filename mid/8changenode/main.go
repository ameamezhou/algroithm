package main

import (
	"fmt"
	"time"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
    Val int
    Next *ListNode
}
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
	var rear = new(ListNode)

	flag := 0
	p := head
	temp := head
	for p != nil && p.Next != nil {
	    rear = p
	    p = p.Next
        rear.Next = p.Next
        p.Next = rear
		if flag == 1 {
			temp.Next = p
			temp = rear
		}
        if flag == 0 {
            head = p
            flag = 1
        }
        p = rear.Next
    }
    return head
}

func main(){
	var t = new(ListNode)
	t.Val = 1
	t.Next = new(ListNode)
	tm := t
	tm = tm.Next
	tm.Val = 2
	tm.Next = new(ListNode)
	tm = tm.Next
	tm.Val = 3
	tm.Next = new(ListNode)
	tm = tm.Next
	tm.Val = 4
	tm.Next = nil
	t = swapPairs(t)
	for ;t.Next != nil;t = t.Next {
		fmt.Println(t.Val)
	}

}
