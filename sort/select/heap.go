package _select

// 堆排序

// 建大根堆
func buildHeap(nums []int) {
	size := len(nums)
	for i := size / 2; i >= 0; i-- {
		down(nums, i, size)
	}
}

// 插入元素的时候需要使用
func up(nums []int, i int) {
	for {
		// 父结点计算公式！！
		parent := (i - 1) / 2
		// 如果父结点大于子节点了，说明不需要调换，已经满足条件了
		// 如果 i == parent 说明 i = 0 了，没有父结点，也不需要调换
		if i == parent || nums[parent] > nums[i] {
			break
		}
		// 父结点没有子节点i的值大，需要调换
		nums[parent], nums[i] = nums[i], nums[parent]
	}
}

// 初始化的时候需要建堆
// 删除元素的时候需要down：
// 将第一个元素和最后一个元素进行调换, 然后重新 down(nums, 0, size)
func down(nums []int, i int, size int) {
	left, right := 2*i+1, 2*i+2
	j := i

	// 找到最大的元素的下标
	if left < size && nums[left] > nums[j] {
		j = left
	}
	if right < size && nums[right] > nums[j] {
		j = right
	}

	if j != i {
		nums[i], nums[j] = nums[j], nums[i]
		down(nums, j, size)
	}
}

// 不用递归也可以
func down2(nums []int, i, size int) {
	for {
		left, right := i*2+1, i*2+2
		// 说明i是叶子节点
		if left >= size {
			break
		}
		// 找到最大的节点
		largest := i
		if left < size && nums[left] > nums[largest] {
			largest = left
		}
		if right < size && nums[right] > nums[largest] {
			largest = right
		}

		if largest != i {
			nums[i], nums[largest] = nums[largest], nums[i]
			i = largest
		} else {
			// 如果最大的节点就是本身，那么后面也不需要进行操作了
			break
		}
	}
}

// 从堆中弹出一个数
// 首先将堆顶的元素和最后一个元素进行调换，然后进行down操作
func pop(arr *[]int) int {
	nums := *arr
	n := len(nums) - 1
	nums[0], nums[n] = nums[n], nums[0]
	down(nums, 0, n)
	x := nums[n]

	*arr = (*arr)[:n]
	return x
}

func HeapSort(nums []int) {
	size := len(nums)
	// 首先建堆
	buildHeap(nums)
	// 然后这个时候最大的元素在第一个
	for size > 1 {
		// 第一个元素和最后一个进行交换
		nums[0], nums[size-1] = nums[size-1], nums[0]
		size--
		down(nums, 0, size)
	}
}