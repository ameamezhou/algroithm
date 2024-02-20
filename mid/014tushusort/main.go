package main

import (
	"fmt"
)

// https://leetcode.cn/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/

func validateBookSequences1(putIn []int, takeOut []int) bool {
	// 用栈
	n := len(putIn)
	stack := make([]int, n)
	sit := 0
	iIn := 0
	iOut := 0
	result := 0
	count := 0
	for {
		if count >= 2 * n {
			break
		}
		fmt.Println(stack)
		if sit == 0 && iIn < n - 1 {
			stack[sit] = putIn[iIn]
			sit++
			iIn++
		}

		if stack[sit-1] == takeOut[iOut] {
			result++
			iOut++
			sit--
		} else {
			if iIn < n - 1 {
				stack[sit] = putIn[iIn]
				sit++
				iIn++
			}
		}
		count++
	}

	return result == n
}


func validateBookSequences(putIn []int, takeOut []int) bool {
	// 用栈
	n := len(putIn)
	stack := make([]int, n)
	sit := 0
	index := 0
	for _, v := range putIn {
		stack[sit] = v
		sit++
		for sit > 0 && stack[sit-1] == takeOut[index] {
			sit--
			index++
		}
	}
	return index == n
}
func main(){
	p := []int{6,7,8,9,10,11}
	t := []int{9,11,10,8,7,6}
	validateBookSequences(p, t)
}
