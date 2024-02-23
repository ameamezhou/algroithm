package main

import "fmt"

func uniquePaths(m int, n int) int {
	count := 0
	var countFunc func(x, y int)

	countFunc = func(x, y int) {
		if x == m && y == n {
			count++
		} else {
			if x < m {
				countFunc(x+1, y)
			}
			if y < n {
				countFunc(x, y+1)
			}
		}
	}
	countFunc(1, 1)
	return count
}

func uniquePath(m int, n int) int {
	var countSlice = make([][]int, m)
	for i := 0; i < m; i++ {
		countSlice[i] = make([]int, n)
		countSlice[i][0] = 1
	}
	for j := 1; j < n; j++ {
		countSlice[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			countSlice[i][j] = countSlice[i-1][j] + countSlice[i][j-1]
		}
	}
	return countSlice[m-1][n-1]
}

func main(){
	fmt.Println(uniquePaths(3, 3))
}