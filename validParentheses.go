package main

import (
	"fmt"
)

func isValid(s string) bool {
	var stack []rune
	var openBrackets = map[rune]rune{'(': ')', '{': '}', '[': ']'}
	for _, v := range s {
		if _, ok := openBrackets[v]; ok {
			stack = append(stack, v)
		} else {
			// v - закрывающая скобка
			if len(stack) == 0 || openBrackets[stack[len(stack)-1]] != v {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	if len(stack) == 0 {
		return true
	}

	return false
}

func main() {
	var s = "()[]{}"
	fmt.Println(isValid(s))

}
