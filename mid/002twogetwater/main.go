package main


func maxArea(height []int) int {
	head := 0
	rear := len(height) - 1
	var result = 0
	var temp = 0
	for ;head < rear; {
		if height[head] < height[rear] {
			temp = (rear - head) * height[head]
			head += 1
		} else {
			temp = (rear - head) * height[rear]
			rear -= 1
		}
		if temp > result {
			result = temp
		}
	}
	return result
}