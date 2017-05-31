/*
The count-and-say sequence is the sequence of integers beginning as follows:
1, 11, 21, 1211, 111221, ...

1 is read off as "one 1" or 11.
11 is read off as "two 1s" or 21.
21 is read off as "one 2, then one 1" or 1211.
Given an integer n, generate the nth sequence.

Note: The sequence of integers will be represented as a string.
*/

package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func count(s string) string {
	var buffer bytes.Buffer
	num := s[0]
	repeatNums := 1
	for i := 1; i < len(s); i++ {
		if s[i] == num {
			repeatNums++
		} else {
			buffer.WriteString(strconv.Itoa(repeatNums))
			buffer.WriteByte(num)
			num = s[i]
			repeatNums = 1
		}
	}
	buffer.WriteString(strconv.Itoa(repeatNums))
	buffer.WriteByte(num)

	return buffer.String()
}

func countAndSay(n int) string {
	result := "1"
	for i := 1; i < n; i++ {
		result = count(result)
	}
	return result
}
func main() {
	fmt.Println(countAndSay(7))
}
