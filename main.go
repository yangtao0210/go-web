package main

import (
	"strconv"
)

func main() {
	strconv.Itoa(0)
}

func reverseStr(s string, k int) string {
	unit, l, size := 2*k, len(s), len(s)
	left, right, loop, div := 0, k-1, l/unit, l%unit
	bytes := []byte(s)
	for loop > 0 {
		reverseSpecailIndex(bytes, left, right)
		left += unit
		right += unit
		loop--
	}
	if div < k {
		reverseSpecailIndex(bytes, (size/unit)*unit, size-1)
	} else {
		reverseSpecailIndex(bytes, (size/unit)*unit, (size/unit)*unit+k-1)
	}
	return string(bytes)
}

func reverseSpecailIndex(bytes []byte, start, end int) {
	for start < end {
		bytes[start], bytes[end] = bytes[end], bytes[start]
		start++
		end--
	}
}
