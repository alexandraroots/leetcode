package main

import "fmt"

func majorityElement(nums []int) int {
	var counter = make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}

	for num, count := range counter {
		if count > len(nums)/2 {
			return num
		}
	}

	return 0

}

func main() {
	var nums = []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}
