package main

import "fmt"

func removeDuplicates(nums []int) int {
	var nextIdx int
	var alreadySeen = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := alreadySeen[nums[i]]; ok {
			nums[i] = '_'
		} else {
			nums[nextIdx] = nums[i]
			alreadySeen[nums[i]] = 0
			if i != nextIdx {
				nums[i] = '_'
			}
			nextIdx++
		}
	}

	return nextIdx
}

func main() {
	var nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}
