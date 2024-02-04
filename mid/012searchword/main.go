package main

// https://leetcode.cn/problems/word-search/

var station = [][]int{
	{1, 0}, {-1, 0}, {0, 1}, {0, -1},
}

func exist(board [][]byte, word string) bool {
	// bfs
	n := len(word)
	nx := len(board)
	ny := len(board[0])
	flag := make([][]bool, nx)
	for i := range flag {
		flag[i] = make([]bool, ny)
	}
	var check func(x, y, k int)bool
	check = func(x, y, k int) bool {
		if board[x][y] != word[k] {
			return false
		}
		if k == n - 1{
			return true
		}
		flag[x][y] = true
		defer func() {
			flag[x][y] = false
		}()
		for _, sit := range station {
			if newX, newY := x+sit[0], y+sit[1]; 0 <= newX && newX < nx && 0 <= newY && newY < ny && !flag[newX][newY] {
				if check(newX, newY, k+1) {
					return true
				}
			}
		}
		return false
	}
	for x, tempS := range board{
		for y := range tempS{
			if check(x, y, 0) {
				return true
			}
		}
	}
	return false
}