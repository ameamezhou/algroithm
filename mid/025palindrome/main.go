package main

import "fmt"

// https://leetcode.cn/problems/longest-palindromic-substring/description/


func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	var result string
	var a string
	for pre := 0 ;pre < len(s)-1; pre++ {
		for tail := pre+1; tail <= len(s); tail++ {
			a = s[pre:tail]
			if a == reverse(a) {
				if len(a) > len(result) {
					result = a
				}
			}
			if len(result) == len(s) {
				return result
			}
		}
	}
	return result
}

func reverse(s string) string {
	b := []byte(s)
	for i := 0; i < len(b)/2;i++ {
		j := len(b) - i - 1
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func main(){
	a := "A"
	fmt.Println(a[0:1])
}

func longestPalindrome2(s string) string {
	//双指针法
	start := 0//记录最长回文子串的起始位置
	end := 0 //记录最长回文子串的末尾位置
	ln := len(s)//s的长度
	for i := 0; i < ln; {
		//左右指针，向两边扩展
		l, r := i, i

		//先考虑回文子串可能是偶数的情况，让r先走
		for r < ln -1 && s[r] == s[r+1] { //防止越界，一直向后比较，不相等停止
			r++
		}
		//i到达r所扩展的最远长度的下一个字符
		i = r+1

		//两边一起扩展
		for l > 0 && r < ln-1 && s[l-1] == s[r+1] {
			l--
			r++
		}

		//更新最长回文子串的长度和起始点
		if end -start < r - l {
			start = l
			end = r
		}

	}

	return s[start : end+1]
}