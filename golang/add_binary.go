package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	aLen := len(a)
	bLen := len(b)
	index := aLen
	if aLen < bLen {
		index = bLen
	}
	result := make([]byte, index+1)
	var c uint8
	aLen--
	bLen--
	for aLen >= 0 && bLen >= 0 {
		tmp := a[aLen] - "0"[0] + b[bLen] - "0"[0] + c
		result[index] = tmp % 2
		if tmp >= 2 {
			c = 1
		} else {
			c = 0
		}

		index--
		aLen--
		bLen--
	}

	for aLen >= 0 {
		tmp := a[aLen] - "0"[0] + c
		result[index] = tmp % 2
		if tmp >= 2 {
			c = 1
		} else {
			c = 0
		}

		aLen--
		index--
	}

	for bLen >= 0 {
		tmp := b[bLen] - "0"[0] + c
		result[index] = tmp % 2
		if tmp >= 2 {
			c = 1
		} else {
			c = 0
		}

		bLen--
		index--
	}
	if c > 0 {
		result[index] = c
	}

	test := ""
	first := true
	for _, r := range result {
		if r == 0 && first {
			continue
		} else {
			first = false
		}
		test += strconv.Itoa(int(r))

	}
	return test

}

func main() {
	a := "11"
	b := "1"
	fmt.Println(addBinary(a, b))

}
