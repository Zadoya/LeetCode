// условие задачи - https://leetcode.com/problems/simplify-path/description

package main

import (
	"strings"
)

func split(str string, delimiter byte) []string {
	dirs := make([]string, 0)
	for i := 0; i < len(str); {
		if str[i] == delimiter {
			i++
			dir := make([]byte, 0)
			for i < len(str) && str[i] != delimiter {
				dir = append(dir, str[i])
				i++
			}
			if len(dir) > 0 && string(dir) != "." {
				if string(dir) == ".." {
					if len(dirs) > 0 {
						dirs = dirs[:len(dirs) - 1]
					}
				} else {
					dirs = append(dirs, string(dir))
				}
			}
		} else {
			i++
		}
	}

	return dirs
}

func simplifyPath(path string) string {
	var result strings.Builder
	dirs := split(path, '/')
	for i := range dirs {
		result.WriteString("/" + dirs[i])
	}
	if result.Len() == 0 {
		return "/"
	}
	return result.String()
}
