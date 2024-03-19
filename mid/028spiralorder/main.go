package main

import "fmt"

// https://leetcode.cn/problems/spiral-matrix/description/?envType=study-plan-v2&envId=top-100-liked

func spiralOrder(matrix [][]int) []int {
	x, y := 0, 0
	xLen := len(matrix[0]) - 1
	yLen := len(matrix) - 1
	xSta := 0
	ySta := 0
	var result = []int{}
	for xSta <= xLen && ySta <= yLen {
		for ; x < xLen; x++ {
			result = append(result, matrix[y][x])
			fmt.Println(result, 1, x, y)
		}
		ySta++
		y++
		for ; y < yLen; y++ {
			result = append(result, matrix[y][x])
			fmt.Println(result, 2, x, y)
		}
		xLen--
		x--
		for ; x > xSta; x-- {
			result = append(result, matrix[y][x])
			fmt.Println(result, 3, x, y)
		}
		yLen--
		y--
		for ; y > ySta; y-- {
			result = append(result, matrix[y][x])
			fmt.Println(result, 4, x, y)
		}
		xSta++
		x++
	}
	return result
}

func main(){
	matrix := [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	fmt.Println(spiralOrder(matrix))
}

func spiralOrder2(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var (
		rows, columns = len(matrix), len(matrix[0])
		order = make([]int, rows * columns)
		index = 0
		left, right, top, bottom = 0, columns - 1, 0, rows - 1
	)

	for left <= right && top <= bottom {
		for column := left; column <= right; column++ {
			order[index] = matrix[top][column]
			index++
		}
		for row := top + 1; row <= bottom; row++ {
			order[index] = matrix[row][right]
			index++
		}
		if left < right && top < bottom {
			for column := right - 1; column > left; column-- {
				order[index] = matrix[bottom][column]
				index++
			}
			for row := bottom; row > top; row-- {
				order[index] = matrix[row][left]
				index++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return order
}