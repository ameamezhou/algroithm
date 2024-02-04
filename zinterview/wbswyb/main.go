package main

import "math"

/*
两个单向链表
如果有相交的结点  Y

如果没有相交就是平行线

找到相交的结点
*/

type Node struct {
	value int
	next  *Node
}

func Check(n1, n2 *Node)bool{
	count1 := 0
	count2 := 0
	p1 := n1
	p2 := n2
	for p1 != nil {
		p1 = p1.next
		count1++
	}

	for p2 != nil {
		p2 = p2.next
		count2++
	}

	for i := 0; i < int(math.Abs(float64(count1 - count2))); i++ {
		if count1 > count2 {
			n1 = n1.next
		} else {
			n2 = n2.next
		}
	}

	for n1 != nil {
		if n1.value == n2.value {
			return true
		}
		n1 = n1.next
		n2 = n2.next
	}

	return false
}
