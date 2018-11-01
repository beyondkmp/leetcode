package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	var tmp []rune
	var result []string

	for _, s := range path + "/" {
		if s == '/' {
			switch string(tmp) {
			case "":
			case ".":
			case "..":
				if len(result) > 0 {
					result = result[:len(result)-1]
				}
			default:
				result = append(result, string(tmp))
			}
			tmp = nil
		} else {
			tmp = append(tmp, s)
		}
	}
	return "/" + strings.Join(result, "/")
}

func main() {
	fmt.Println(simplifyPath("/a/./b/../../c/"))
	fmt.Println(simplifyPath("/..."))
	fmt.Println(simplifyPath("/../"))
}
