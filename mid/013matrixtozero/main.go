package main

func setZeroes(matrix [][]int)  {
	y := len(matrix)
	x := len(matrix[0])
	var flag [][]int8
	flag = make([][]int8, y)
	for i := 0; i < y; i++ {
		flag[i] = make([]int8, x)
	}

	var change func(i, j int)
	change = func(i, j int) {
		flag[i][j] += 1
		for k := range matrix[i] {
			if flag[i][k] > 0 || matrix[i][k] == 0 {
				continue
			}
			flag[i][k] += 1
			matrix[i][k] = 0
		}
		for v := range matrix {
			if flag[v][j] > 0 || matrix[v][j] == 0 {
				continue
			}
			flag[v][j] += 1
			matrix[v][j] = 0
		}
	}
	for i := 0 ; i < y; i++ {
		for j := 0 ; j < x; j++ {
			if matrix[i][j] == 0 && flag[i][j] == 0{
				change(i, j)
			}
		}
	}
}
