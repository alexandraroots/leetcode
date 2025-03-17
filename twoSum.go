package main

import "fmt"

func twoSum(nums []int, target int) []int {
	tmpMap := make(map[int]int)

	for idx1, val := range nums {
		needed := target - val

		if idx2, ok := tmpMap[needed]; ok {
			return []int{idx1, idx2}
		}

		tmpMap[val] = idx1
	}
	return nil
}

func main() {
	nums := []int{3, 3}
	target := 6
	r := twoSum(nums, target)
	fmt.Println(r)

}
