package main

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	if n <= 6 {
		return true
	}
	factor := []int{2, 3, 5}
	for _, v := range factor {
		for n%v == 0 {
			n /= v
		}
	}

	return n == 1
}