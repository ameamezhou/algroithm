package main

import "fmt"

// https://leetcode.cn/problems/jump-game-iii/


func canReach(arr []int, start int) bool {
	l := len(arr)
	flag := make([]bool, l)
	flag[start] = true

	var tempFunc func(s int) bool
	tempFunc = func(s int) bool {
		if arr[s] == 0 {
			return true
		}
		flag[s] = true
		if s - arr[s] >= 0 && !flag[s - arr[s]] {
			if tempFunc(s - arr[s]) {
				return true
			}
		}

		if s + arr[s] <= l - 1 && !flag[s + arr[s]] {
			if tempFunc(s + arr[s]) {
				return true
			}
		}
		return false
	}

	return tempFunc(start)
}


func canReach2(arr []int, start int) bool {
	result := false
	var tempFunc func(s int)
	l := len(arr)
	var tempMap = map[int]byte{}
	tempFunc = func(s int) {
		if s < 0 || s >= l {
			return
		}
		if _, ok := tempMap[s]; ok {
			return
		}
		tempMap[s] = 1

		if arr[s] == 0 || result {
			result = true
			return
		}
		tempFunc(s + arr[s])
		tempFunc(s - arr[s])
	}

	tempFunc(start)
	return result
}

func main() {
	arr := []int{4, 2, 3, 0, 3, 1, 2}
	fmt.Println(canReach(arr, 0))
}