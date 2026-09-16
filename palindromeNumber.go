package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var inverseX int
	var copyX = x

	for x > 0 {
		inverseX *= 10
		inverseX += x % 10
		x = x / 10
	}

	return copyX == inverseX

}

func main() {
	var x = 121
	fmt.Println(isPalindrome(x))

}
