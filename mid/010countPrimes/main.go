package main

import "fmt"

// 暴力解法
func countPrimes2(n int) int {
	if n <= 2 {
		return 0
	}
	count := 0
	for i := 2; i <= n - 1; i++ {
		for j := 2; j <= i; j++ {
			if i % j == 0 && i != j {
				break
			}
			if j >= i/2 || i == j {
				fmt.Println(i, j)
				count++
				break
			}
		}
	}
	return count
}

// 筛选法
func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	flag := make([]bool, n)
	count := 0
	for i := 2; i <= n - 1; i++ {
		if !flag[i] {
			count++
		}
		if flag[i] {
			continue
		}
		for j := i*i; j < n; j += i{
			flag[j] = true
		}
	}
	return count
}

func main(){
	n := 3
	countPrimes(n)
}
