package main

import "fmt"

// https://leetcode.cn/problems/trapping-rain-water/

func trap(height []int) int {
	if len(height) < 1 {
		return 0
	}
	result := 0
	j := 0
	pre := height[0]
	temp := 0
	var rear int
	for i := 1; i < len(height); i++ {
		fmt.Println(pre, rear)
		if i - j == 1 && height[i] >= pre {
				pre = height[i]
				j = i
				continue
		}
		if height[i] >= pre {
			rear = height[i]
			result = result + min(pre, rear) * (i - j - 1) - temp
			pre = rear
			j = i
			temp = 0
			continue
		}
		temp += height[i]
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main(){
	temSlice := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	trap(temSlice)
}

func trap1(height []int) int {
	var result int
	stack := []int{}
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			curWidth := i - left - 1
			curHeight := min(height[left], h) - height[top]
			result += curWidth * curHeight
		}
		stack = append(stack, i)
	}
	return result
}