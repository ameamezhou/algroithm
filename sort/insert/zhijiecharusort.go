package insert

// 直接插入排序

func insertSort(data []int) {
	for i := 1; i < len(data); i++ {
		temp, j := data[i], i
		if temp < data[j-1] {
			for j >= 1 && data[j - 1] > temp {
				data[j] = data[j - 1]
				j--
			}
		}
		data[j] = temp
	}
}
