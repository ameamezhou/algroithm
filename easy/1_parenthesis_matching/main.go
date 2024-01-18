package main

import "fmt"

// https://leetcode.cn/problems/valid-parentheses/description/

/*
示例 1：

输入：s = "()"
输出：true
示例 2：

输入：s = "()[]{}"
输出：true
示例 3：

输入：s = "(]"
输出：false

很简单的思路  栈匹配就好   这里我们用go结构体写结构体方法实现栈结构  增加难度
*/

type StackData struct { // 栈 元素
	dataTypeS 	string // 用来记录字符  其实在go里单个字符可以用byte来记录，用string比较浪费空间
	//dataTypeB 	byte // 用两种方法一起来记录
}

var Check = map[byte]byte{
	')' : '(',
	']' : '[',
	'}' : '{',
}

type Stack struct {	// 栈
	StackSlice	[]byte
	StackLength int 	// 栈长度
}

func (s *Stack) Init (n int){
	s.StackLength = 0
	s.StackSlice = make([]byte, n)
}

func isValid(s string) bool {
	n := len(s)
	if n % 2 != 0 {
		return false
	}
	var stack Stack
	stack.Init(n)

	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack.StackSlice[stack.StackLength] = byte(v)
			stack.StackLength++
		}
		if v == ')' || v == ']' || v == '}' {
			if stack.StackLength == 0 || Check[byte(v)] != stack.StackSlice[stack.StackLength-1] {
				return false
			}
			stack.StackLength--
		}
	}
	if stack.StackLength != 0 {
		return false
	}
	return true
}

// 回文数
/*
https://leetcode.cn/problems/palindrome-number/description/
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

例如，121 是回文，而 123 不是。


示例 1：

输入：x = 121
输出：true
示例 2：

输入：x = -121
输出：false
解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3：

输入：x = 10
输出：false
解释：从右向左读, 为 01 。因此它不是一个回文数。
*/

func IsPalindrome(x int) bool {
	var r int
	r = 0
	for i := x; i > 0; i /= 10 {
		r += i%10
		r *= 10
	}
	r /= 10
	if x == r {
		return true
	}
	return false
}

func main(){
	var a int
	a = 1234567
	fmt.Println(IsPalindrome(a))

}
