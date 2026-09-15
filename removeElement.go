package main

import "fmt"

func removeElement(nums []int, val int) int {
	var counter int
	var deletedIdxs []int
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums[i] = '_'
			deletedIdxs = append(deletedIdxs, i)
		} else {
			counter++
			if len(deletedIdxs) != 0 {
				nums[deletedIdxs[0]] = nums[i]
				nums[i] = '_'
				deletedIdxs = append(deletedIdxs[1:], i)
			}
		}
	}

	return counter
}

func main() {
	var nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	var val = 2
	fmt.Println(removeElement(nums, val))
}
