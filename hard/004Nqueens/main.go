package _04Nqueens

// https://leetcode.cn/problems/n-queens/description/

// 数据结构经典算法题



func solveNQueens(n int) [][]string {
	var result = [][]string{}
	var queens = make([]int, n)
	var flag = make([]bool, n)
	var backTrace func(row int)
	d1 := map[int]bool{}
	d2 := map[int]bool{}
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	backTrace = func(row int) {
		if row == n {
			temp := makeString(queens, n)
			result = append(result, temp)
			return
		}
		for i := 0; i < n; i++ {
			if flag[i] {
				continue
			}
			diagonal1 := row - i
			if d1[diagonal1]{
				continue
			}
			diagonal2 := row + i
			if d2[diagonal2]{
				continue
			}
			flag[i] = true
			d1[diagonal1] = true
			d2[diagonal2] = true
			queens[row] = i
			backTrace(row+1)
			queens[row] = -1
			flag[i] = false
			delete(d1, i+1)
			delete(d2, i-1)
		}
	}
	backTrace(0)
	return result
}

func makeString(queens []int, n int)[]string{
	temp := []string{}
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n;j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		temp = append(temp, string(row))
	}
	return temp
}