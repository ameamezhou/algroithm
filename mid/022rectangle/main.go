package _22rectangle

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/minimum-area-rectangle-ii/description/


func minAreaFreeRect(points [][]int) float64 {
	if len(points) < 4 {
		return 0
	}
	n := len(points)
	m := make(map[string]struct{})
	result := float64(0)
	for _, v := range points {
		s := fmt.Sprintf("%d,%d", v[0], v[1])
		m[s] = struct{}{}
	}

	for i := range points {
		for j := i+1; j < n; j++ {
			for k := j+1; k < n ; k++ {
				r := points[i]
				t1 := points[j]
				t2 := points[k]

				afa1 := []int{t1[0]-r[0], t1[1]-r[1]}
				afa2 := []int{t2[0]-r[0], t2[1]-r[1]}

				if multAfa(afa1, afa2) != 0 {
					continue
				}

				d := fmt.Sprintf("%d,%d", r[0] + afa1[0] + afa2[0], r[1] + afa1[1] + afa2[1])
				if _, ok := m[d]; !ok {
					continue
				}
				s := area(r, t1, t2)
				if s < result || result == 0 {
					result = s
				}
			}
		}
	}
	return result
}

func multAfa(a, b []int) int {
	return a[0]*b[0] + a[1]*b[1]
}

func area(x1, x2, x3 []int) float64 {
	L1 := math.Pow(math.Pow(float64(x3[0]-x1[0]), 2) + math.Pow(float64(x3[1]-x1[1]), 2), 0.5)
	L2 := math.Pow(math.Pow(float64(x2[0]-x1[0]), 2) + math.Pow(float64(x2[1]-x1[1]), 2), 0.5)
	return L1*L2
}