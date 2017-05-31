#  First Unique Character in a String

Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.

**Examples**:

```
s = "leetcode"
return 0.

s = "loveleetcode",
return 2.
```

**Note**: You may assume the string contain only lowercase letters.

## 思路

先记录下所有字符出现的次数，再对所有的字符串重新遍历一次，第一个次数为1的就是**First Unique Character**. 自己一开始也想到了先记录下所有字符出现的次数，但是后面再循环一次就没有想到。这个解法的算法复杂度是O(2n).

```
package main

import "fmt"

func firstUniqChar(s string) int {
	var nums [26]int
	for _, w := range s {
		nums[w-'a'] += 1
	}
	for i, w := range s {
		if nums[w-'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	s := "leetcode"
	fmt.Println(firstUniqChar(s))

	s = "loveleetcode"
	fmt.Println(firstUniqChar(s))
}
```


