package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	var res []byte
	var j = 0
	for i := 0; i < len(strs[0]); i++ {
		var char = strs[0][i]
		for _, str2 := range strs {
			if j == len(str2) || char != str2[j] {
				return string(res)
			}
		}

		res = append(res, char)
		j += 1
	}
	return string(res)
}

func main() {
	strs := []string{"flower", "flower", "flower", "flower"}
	fmt.Println(longestCommonPrefix(strs))

}
