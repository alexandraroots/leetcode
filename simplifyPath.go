package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	var folders = strings.Split(path, "/")
	var rightPath []string
	for _, folder := range folders {
		if folder == "" || folder == "." {
			continue
		}

		if folder == ".." {
			if len(rightPath) > 0 {
				rightPath = rightPath[:len(rightPath)-1]
			}
			continue
		}

		rightPath = append(rightPath, folder)

	}

	return "/" + strings.Join(rightPath, "/")

}

func main() {
	strs := "/a/./b/../../c/"
	fmt.Println(simplifyPath(strs))

}
