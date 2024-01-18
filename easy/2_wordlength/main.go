package main

import "fmt"

func lengthOfLastWord(s string) int {
	var result int
	var flag int
	result = 0
	flag = 0
	for _, v := range s {
		if v == 32 {
			flag = 0
		} else {
			if flag == 1 {
				result++
			}else {
				result = 1
				flag = 1
			}

		}
	}
	return result
}

func main(){
	a := "luffy is still joyboy"
	fmt.Println(lengthOfLastWord(a))
}
