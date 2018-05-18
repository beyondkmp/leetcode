package main

import (
	"fmt"
	"strconv"
)

func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return []string{}
	}

	var res []string
	var dfs func(sub string, index int, res_ip string)

	dfs = func(sub string, index int, res_ip string) {
		if index > 4 {
			return
		}

		if index == 4 && len(sub) == 0 {
			res = append(res, res_ip[:len(res_ip)-1])
		}

		for i := 1; i < 4; i++ {
			if i > len(sub) {
				break
			}

			if string(sub[0]) == "0" && i > 1 {
				break
			}

			v, err := strconv.Atoi(sub[:i])
			if err != nil {
				continue
			}

			if v <= 255 {
				dfs(sub[i:], index+1, res_ip+sub[:i]+".")
			}
		}
	}

	dfs(s, 0, "")
	return res
}

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("192168011"))
	// [255.255.11.135 255.255.111.35]
	// [19.216.80.11 192.16.80.11 192.168.0.11]
}
