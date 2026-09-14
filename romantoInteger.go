package main

import "fmt"

func romanToInt(s string) int {
	var romanDict = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var total int

	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			total += romanDict[s[i]]
		} else {
			var currElem = s[i]
			var nextElem = s[i+1]
			if romanDict[currElem] < romanDict[nextElem] {
				total -= romanDict[currElem]
			} else {
				total += romanDict[currElem]
			}
		}

	}

	return total

}

func main() {
	var s = "MCMXCIV"
	fmt.Println(romanToInt(s))
}
