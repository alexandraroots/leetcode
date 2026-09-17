package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	var splittedString = strings.Split(strings.TrimSpace(s), " ")
	return len(splittedString[len(splittedString)-1])

}

func main() {
	strs := "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(strs))

}
