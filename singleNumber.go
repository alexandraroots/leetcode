package main

import "fmt"

func singleNumber(nums []int) int {
	var doubles = make(map[int]int)
	for _, value := range nums {
		if _, ok := doubles[value]; ok {
			doubles[value]++
		} else {
			doubles[value] = 1
		}
	}

	fmt.Println(doubles)

	var res = 0
	for key, value := range doubles {
		if value == 1 {
			res = key
			break
		}
	}
	return res
}

//func main() {
//	var r int
//	var nums = []int{2, 2, 1}
//	r = singleNumber(nums)
//	fmt.Println(r)
//}
