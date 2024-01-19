package main

import (
	"fmt"
	"sort"
)

func quickSort(nums []int, left, right int) {
	if left >= right{
		return
	}
	i, j := left, right
	// i从左边开始遍历，j从右边开始遍历。
	p := nums[i]
	// 设置基准数
	for i < j{
		for i < j && nums[i] < p{
			// 寻找一个比基准数大的数
			i++
		}
		for i < j && nums[j] > p{
			// 寻找一个比基准数小的数
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
		// 替换这两个数
	}
	nums[j] = p
	// 如果i和j相遇，就将基准数和j下标的值进行替换
	fmt.Println(nums)
	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}

func threeSum(nums []int) [][]int {
	n := len(nums)
	var result = [][]int{}
	if n < 3 {
		return result
	}
	sort.Ints(nums)
	var left, right int
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return result
		}
		left = i + 1
		right = n - 1
		if i > 0 {
			if nums[i] == nums[i - 1] {
				continue
			}
		}
		for left < right {
			if nums[i] + nums[left] + nums[right] > 0 {
				right--
				for nums[right] == nums[right+1] && left < right {
					right--
				}
			} else if nums[i] + nums[left] + nums[right] < 0 {
				left++
				for nums[left] == nums[left-1] && left < right {
					left++
				}
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				for nums[left] == nums[left-1] && left < right {
					left++
				}
			}
		}
	}
	return result
}

func main(){
	nums := []int{0,0,0,0}
	fmt.Println(threeSum(nums))
}
