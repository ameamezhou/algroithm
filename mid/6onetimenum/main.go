package main


func singleNumber(nums []int) []int {
	var result []int
	var tempCache = make(map[int]int8)
	var tempTwice = make(map[int]bool)

	for _, v := range nums {
		if _, ok := tempTwice[v]; ok {
			continue
		}
		if _, ok := tempCache[v]; !ok {
			tempCache[v] = 1
		} else {
			tempTwice[v] = true
			delete(tempCache, v)
		}

	}
	for k, _ := range tempCache {
		result = append(result, k)
	}
	return result
}

func main(){

}