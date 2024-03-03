package main

// https://leetcode.cn/problems/longest-absolute-file-path/

func lengthLongestPath(input string) int {
	result := 0
	pathLength := []int{}
	index := 0
	n := len(input)
	for index < n {
		// 找到 \t 并判断是第几层
		curLevel := 0
		for input[index] == '\t' && index < n {
			curLevel++
			index++
		}

		indexNext := index + 1
		// 找到 \n 代表这一part结束
		fileFlag := false
		for indexNext < n && input[indexNext] != '\n' {
			if input[indexNext] == '.' {
				fileFlag = true
			}
			indexNext++
		}
		index++ // 跳过换行符

		for len(pathLength) >= curLevel {
			pathLength = pathLength[:len(pathLength)-1]
		}

		if len(pathLength) > 0 {
			indexNext += pathLength[len(pathLength)-1] + 1
		}

		// 将记录的长度进行写入
		if fileFlag {
			result = max(result, indexNext)
		} else {
			pathLength = append(pathLength, indexNext)
		}
	}

	return result
}

func max(a, b int)int{
	if a > b{
		return a
	}
	return b
}

func lengthLongestPath1(input string) int {
	result := 0
	pathLength := []int{0}
	index := 0
	n := len(input)
	for index < n {
		// 找到 \t 并判断是第几层
		curLevel := 0
		for input[index] == '\t' && index < n {
			curLevel++
			index++
		}

		indexNext := index + 1
		// 找到 \n 代表这一part结束
		fileFlag := false
		for indexNext < n && input[indexNext] != '\n' {
			if input[indexNext] == '.' {
				fileFlag = true
			}
			indexNext++
		}

		// 将记录的长度进行写入
		if fileFlag {
			// 如果是文件  直接输出具体长度
			result = max(pathLength[curLevel] + indexNext - index, result)
		} else {
			// 记录路径长度
			if len(pathLength) < curLevel+1 {
				pathLength = append(pathLength, pathLength[curLevel] + indexNext - index + 1)
			} else {
				pathLength[curLevel] = indexNext - index + 1
			}
		}
		index = indexNext
	}

	return result
}

func max1(a, b int)int{
	if a > b{
		return a
	}
	return b
}