package main

import "fmt"

const (
	bucketCntBits = 3
	bucketCnt     = 1 << bucketCntBits
)

func main(){
	fmt.Println(bucketCntBits)
	fmt.Println(bucketCnt)
}
