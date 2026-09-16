package main

import "fmt"

func mySqrt(x int) int {
	var left = 1
	var right = x
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == x {
			return mid
		}

		if mid*mid < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left - 1
}

func main() {
	var x = 4
	fmt.Println(mySqrt(x))

}
