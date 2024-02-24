package uniquepath

import "fmt"

// https://leetcode.cn/problems/unique-paths-ii/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n, m := len(obstacleGrid), len(obstacleGrid[0])
	f := make([]int, m)
	if obstacleGrid[0][0] == 0 {
		f[0] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				f[j] = 0
				continue
			}
			if j - 1 >= 0 && obstacleGrid[i][j-1] == 0 {
				f[j] += f[j-1]
			}
		}
	}
	return f[len(f)-1]
}

func uniquePathsWithObstacles1(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	var countSlice = make([][]int, m)
	flag := 0
	for i := 0; i < m; i++ {
		countSlice[i] = make([]int, n)
		if obstacleGrid[i][0] == 1 {
			flag = 1
		}
		if flag == 0 {
			countSlice[i][0] = 1
		}
	}
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 1 || countSlice[0][0] == 0 {
			break
		}
		countSlice[0][j] = 1
	}
	fmt.Println(countSlice)
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			fmt.Println(countSlice)
			if obstacleGrid[i][j] != 1 {
				countSlice[i][j] = countSlice[i-1][j] + countSlice[i][j-1]
			}
		}
	}
	return countSlice[m-1][n-1]
}
