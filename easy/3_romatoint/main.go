package main

import "fmt"

// https://leetcode.cn/problems/roman-to-integer/description/

var valueMap = map[int32]int {
	'I' : 1,
	'V' : 5,
	'X' : 10,
	'L' : 50,
	'C' : 100,
	'D' : 500,
	'M' : 1000,
}

func romanToInt(s string) int {
	var pre int32 = 0
	var result = 0
	var flag = 0
	for _, v := range s {
		if flag == 0 {
			if pre != 0 {
				if (pre == 'I' && (v == 'V' || v == 'X')) || (pre == 'X' && (v == 'L' || v == 'C')) || (pre == 'C' && (v == 'D' || v == 'M')) {
					result += valueMap[v] - valueMap[pre]

					flag = 1
				} else {
					result += valueMap[pre]
				}
			}
			pre = v
		} else {
			pre = v
			flag = 0
		}
	}
	if flag == 0 {
		result += valueMap[pre]
	}
	return result
}

func main(){
	s := "III"
	fmt.Println(romanToInt(s))
}