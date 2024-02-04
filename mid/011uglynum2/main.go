package main

import "fmt"

/*
给你一个整数 n ，请你找出并返回第 n 个 丑数 。

丑数 就是质因子只包含 2、3 和 5 的正整数。



示例 1：

输入：n = 10
输出：12
解释：[1, 2, 3, 4, 5, 6, 8, 9, 10, 12] 是由前 10 个丑数组成的序列。
示例 2：

输入：n = 1
输出：1
解释：1 通常被视为丑数。
*/

func nthUglyNumber(n int) int {
	if n <= 6 {
		return n
	}
	tempSlice := make([]int, n)
	flag2 := 0
	flag3 := 0
	flag5 := 0
	tempSlice[0] = 1

	for i := 1; i < n; i++ {
		temp2 := tempSlice[flag2] * 2
		temp3 := tempSlice[flag3] * 3
		temp5 := tempSlice[flag5] * 5
		fmt.Println(temp2, temp3, temp5)
		if temp2 <= temp3 && temp2 <= temp5 {
			tempSlice[i] = temp2
			flag2++
		}
		if temp3 <= temp2 && temp3 <= temp5 {
			tempSlice[i] = temp3
			flag3++
		}
		if temp5 <= temp3 && temp5 <= temp2 {
			tempSlice[i] = temp5
			flag5++
		}
		fmt.Println(tempSlice)
	}
	return tempSlice[n-1]
}

func main(){
	nthUglyNumber(10)
}