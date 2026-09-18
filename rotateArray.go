package main

import "fmt"

func rotate(nums []int, k int) {
	var r = k % len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, r-1)
	reverse(nums, r, len(nums)-1)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7}
	var k = 3
	rotate(nums, k)
	fmt.Println(nums)
}
