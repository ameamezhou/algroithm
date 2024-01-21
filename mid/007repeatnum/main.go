package main


func findDuplicate(nums []int) int {
	var mCache = make(map[int]bool)

	for _, v := range nums {
		if _, ok := mCache[v]; ok {
			return v
		}
		mCache[v] = true
	}
	return 0
}

func main(){

}
