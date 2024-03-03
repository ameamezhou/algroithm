package main

import "fmt"

/*
给你一个大小为 n 的字符串数组 strs ，其中包含n个字符串 , 编写一个函数来查找字符串数组中的最长公共前缀，返回这个公共前缀
*/

func longestCommonPrefix( strs []string ) string {
	// write code here
	if len(strs) == 0 {
		return ""
	}
	result := strs[0]
	for i := 1 ; i < len(strs); i++ {
		result = compare(result, strs[i])
	}
	return result
}

func compare(a, b string) string{
	n := min(len(a), len(b))
	i := 0
	for i < n {
		if a[i] != b[i] {
			break
		}
		i++
	}
	return a[0:i]
}

func min(a, b int)int {
	if a < b {
		return a
	}
	return b
}

func main()  {
	s := []string{"abca","abc","abca","abc","abcc"}
	fmt.Println(longestCommonPrefix(s))
}