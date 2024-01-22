package main

import (
	"fmt"
	"strconv"
)

func maximumSwap1(num int) int {
	result := 0
	a := num
	tempSlice := []int{}
	for a > 0 {
		tempSlice = append(tempSlice, a%10)
		a /= 10
	}
	maxmum := 0
	index := 0
	for i := len(tempSlice) - 1; i >= 0; i-- {
		if tempSlice[i] > maxmum {
			maxmum = tempSlice[i]
			index = i
		}
	}

	for i := len(tempSlice) - 1; i >= index; i-- {
		if tempSlice[i] < maxmum {
			tempSlice[i], tempSlice[index] = tempSlice[index], tempSlice[i]
			break
		}
	}

	for i := len(tempSlice) - 1; i >= 0; i-- {
		result = result * 10 + tempSlice[i]
	}
	return result
}

func maximumSwap(num int) int {
	s := []byte(strconv.Itoa(num))
	n := len(s)
	d := make([]int, n)
	for i := range d {
		d[i] = i
	}
	for i := n - 2; i >= 0; i-- {
		if s[i] <= s[d[i+1]] {
			d[i] = d[i+1]
		}
	}
	for i, j := range d {
		if s[i] < s[j] {
			s[i], s[j] = s[j], s[i]
			break
		}
	}
	ans, _ := strconv.Atoi(string(s))
	return ans
}

func main() {
	fmt.Println('8')
}
